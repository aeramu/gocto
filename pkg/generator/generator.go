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
	g.generateEntity()
	g.generateService()
	g.generateAPI()
	g.generateTest()
	g.generateInterface()
}

func (g *Generator) generateEntity() {
	var entities []types.Struct
	for _, entityName := range g.EntitiesName {
		entity := types.NewStruct(entityName)
		entities = append(entities, *entity)
	}
	entityBuffer := &buffer.Buffer{}
	header := types.NewHeader("entity")
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

func (g *Generator) generateService() {
	// generate service interface, implementation
	serviceInterface := types.NewInterface("Service")
	var serviceImplementations []types.Function
	for _, methodName := range g.MethodsName {
		method := types.NewMethod(methodName).
			AddParam(types.NewVariable("ctx", "context.Context")).
			AddParam(types.NewVariable("req", "api."+requestType(methodName))).
			AddReturn("*" + "api." + responseType(methodName)).
			AddReturn("error")
		serviceInterface.AddMethod(*method)

		implementFunction := types.NewFunctionFromMethod(method).
			WithReceiver(types.NewVariable("s", "*service")).
			AddStatement("panic(\"implement me\")")
		serviceImplementations = append(serviceImplementations, *implementFunction)
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
	header := types.NewHeader("service")
	header.AddImportedPackage("context")
	header.AddImportedPackage(g.ModulePath + "/service/api")
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

func (g *Generator) generateAPI() {
	var serviceAPI []types.Struct
	var serviceAPIValidation []types.Function
	for _, methodName := range g.MethodsName {
		requestStruct := types.NewStruct(requestType(methodName))
		responseStruct := types.NewStruct(responseType(methodName))
		serviceAPI = append(serviceAPI, *requestStruct, *responseStruct)

		validateFunction := types.NewFunction("Validate").
			WithReceiver(types.NewVariable("req", requestType(methodName))).
			AddReturn("error").
			AddStatement("return nil")
		serviceAPIValidation = append(serviceAPIValidation, *validateFunction)
	}

	apiBuffer := &buffer.Buffer{}
	header := types.NewHeader("api")
	header.Render(apiBuffer)
	for _, renderer := range serviceAPI {
		renderer.Render(apiBuffer)
	}
	for _, renderer := range serviceAPIValidation {
		renderer.Render(apiBuffer)
	}

	apiFile, err := os.Create("service/api/" + "api.go")
	if err != nil {
		panic(err)
	}
	apiBuffer.Flush(apiFile)
}

func (g *Generator) generateInterface() {
	// generate adapter struct
	adapter := types.NewStruct("Adapter")
	var interfaces []types.Interface
	for _, interfaceName := range g.InterfacesName {
		adapter.AddVariable(types.NewVariable(interfaceName, interfaceName))
		interfaces = append(interfaces, *types.NewInterface(interfaceName))
	}

	b := &buffer.Buffer{}
	header := types.NewHeader("service")
	header.Render(b)
	adapter.Render(b)
	for _, renderer := range interfaces {
		renderer.Render(b)
	}

	f, _ := os.Create("service/" + "interface.go")
	b.Flush(f)
}

func requestType(s string) string {
	return fmt.Sprintf("%sReq", s)
}

func responseType(s string) string {
	return fmt.Sprintf("%sRes", s)
}
