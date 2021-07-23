package generator

import (
	"fmt"
	"github.com/aeramu/gocto/pkg/buffer"
	"github.com/aeramu/gocto/pkg/types"
	"os"
)

type Generator struct {
	ModulePath     string
	EntitiesName   []string
	MethodsName    []string
	InterfacesName []string
}

func (g *Generator) Generate() {
	g.generateEntityFile()
	g.generateServiceFile()
	g.generateAPIFile()
	g.generateTestFile()
	g.generateInterfaceFile()
}

func (g *Generator) generateEntityFile() {
	header := types.NewHeader("entity")

	var entities []types.Struct
	for _, entityName := range g.EntitiesName {
		entity := types.NewStruct(entityName)
		entities = append(entities, entity)
	}

	entityBuffer := &buffer.Buffer{}
	header.Render(entityBuffer)
	for _, renderer := range entities {
		renderer.Render(entityBuffer)
	}

	entityFile, err := os.Create("entity/" + "entity.go")
	if err != nil {
		panic(err)
	}
	entityBuffer.Flush(entityFile)
}

func (g *Generator) generateServiceFile() {
	header := types.NewHeader("service").
		AddImportedPackage("context").
		AddImportedPackage(g.ModulePath + "/service/api")

	// generate service interface, implementation
	serviceInterface := types.NewInterface("Service")
	var serviceImplementations []types.Function
	for _, methodName := range g.MethodsName {
		method := types.NewMethod(methodName).
			AddParam(types.NewVariable("ctx", "context.Context")).
			AddParam(types.NewVariable("req", "api."+requestType(methodName))).
			AddReturn("*" + "api." + responseType(methodName)).
			AddReturn("error")
		serviceInterface.AddMethod(method)

		implementFunction := types.NewFunctionFromMethod(method).
			WithReceiver(types.NewVariable("s", "*service")).
			AddStatement("panic(\"implement me\")")
		serviceImplementations = append(serviceImplementations, implementFunction)
	}

	// generate service constructor
	adapter := types.NewVariable("adapter", "Adapter")
	serviceConstructor := types.NewFunction("NewService").
		AddParam(adapter).
		AddReturn("Service").
		AddStatement("return &service {").
		AddStatement("\t%s: %s,", adapter.Name, adapter.Name).
		AddStatement("}")

	// generate service struct
	serviceStruct := types.NewStruct("service").AddVariable(adapter)

	serviceBuffer := &buffer.Buffer{}
	header.Render(serviceBuffer)
	serviceInterface.Render(serviceBuffer)
	serviceConstructor.Render(serviceBuffer)
	serviceStruct.Render(serviceBuffer)
	for _, renderer := range serviceImplementations {
		renderer.Render(serviceBuffer)
	}

	serviceFile, err := os.Create("service/" + "service.go")
	if err != nil {
		panic(err)
	}
	serviceBuffer.Flush(serviceFile)
}

func (g *Generator) generateAPIFile() {
	header := types.NewHeader("api")

	serviceAPI := []types.Struct{}
	serviceAPIValidation := []types.Function{}
	for _, methodName := range g.MethodsName {
		requestStruct := types.NewStruct(requestType(methodName))
		responseStruct := types.NewStruct(responseType(methodName))
		serviceAPI = append(serviceAPI, requestStruct, responseStruct)

		validateFunction := types.NewFunction("Validate").
			WithReceiver(types.NewVariable("req", requestType(methodName))).
			AddReturn("error").
			AddStatement("return nil")
		serviceAPIValidation = append(serviceAPIValidation, validateFunction)
	}

	apiBuffer := &buffer.Buffer{}
	header.Render(apiBuffer)
	for i := range g.MethodsName {
		serviceAPI[i].Render(apiBuffer)
		serviceAPIValidation[i].Render(apiBuffer)
	}

	apiFile, err := os.Create("service/api/" + "api.go")
	if err != nil {
		panic(err)
	}
	apiBuffer.Flush(apiFile)
}

func (g *Generator) generateInterfaceFile() {
	header := types.NewHeader("service")

	// generate adapter struct
	adapter := types.NewStruct("Adapter")
	var interfaces []types.Interface
	for _, interfaceName := range g.InterfacesName {
		adapter.AddVariable(types.NewVariable(interfaceName, interfaceName))
		interfaces = append(interfaces, types.NewInterface(interfaceName))
	}

	b := &buffer.Buffer{}
	header.Render(b)
	adapter.Render(b)
	for _, renderer := range interfaces {
		renderer.Render(b)
	}

	f, _ := os.Create("service/" + "interface.go")
	b.Flush(f)
}

func (g *Generator) generateTestFile() {
	header := types.NewHeader("service").
		AddImportedPackage("context").
		AddImportedPackage("github.com/stretchr/testify/assert").
		AddImportedPackage(g.ModulePath + "/mocks").
		AddImportedPackage(g.ModulePath + "/service/api").
		AddImportedPackage("testing")

	initTest := types.NewFunction("initTest")
	for _, interfaceName := range g.InterfacesName {
		initTest.AddStatement("mock%s = new(mocks.%s)", interfaceName, interfaceName)
	}
	initTest.AddStatement("adapter = Adapter {")
	for _, interfaceName := range g.InterfacesName {
		initTest.AddStatement("\t%s: mock%s,", interfaceName, interfaceName)
	}
	initTest.AddStatement("}")

	var testFunctions []types.Function
	for _, methodName := range g.MethodsName {
		testFunctions = append(testFunctions, testFunction(methodName))
	}

	testBuffer := &buffer.Buffer{}
	header.Render(testBuffer)

	testBuffer.Println("")
	testBuffer.Println("var (")
	testBuffer.Println("\tadapter Adapter")
	for _, interfaceName := range g.InterfacesName {
		testBuffer.Println("\tmock%s *mocks.%s", interfaceName, interfaceName)
	}
	testBuffer.Println(")")

	initTest.Render(testBuffer)
	for _, renderer := range testFunctions {
		renderer.Render(testBuffer)
	}

	testFile, err := os.Create("service/" + "service_test.go")
	if err != nil {
		panic(err)
	}
	testBuffer.Flush(testFile)
}

func requestType(s string) string {
	return fmt.Sprintf("%sReq", s)
}

func responseType(s string) string {
	return fmt.Sprintf("%sRes", s)
}

func testFunction(methodName string) types.Function {
	return types.NewFunction(fmt.Sprintf("Test_%s_%s", "service", methodName)).
		AddParam(types.NewVariable("t", "*testing.T")).
		AddStatement(testTemplate, methodName, methodName, methodName)
}

const (
	testTemplate = `var (

	)
	type args struct {
		ctx context.Context
		req api.%sReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.%sRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.%s(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}`
)