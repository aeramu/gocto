package generator

import (
	"fmt"
	"github.com/aeramu/gocto/pkg/buffer"
	"github.com/aeramu/gocto/pkg/types"
	"os"
)

func (g *Generator) generateTest() {
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
	header := types.NewHeader("service").
		AddImportedPackage("context").
		AddImportedPackage("github.com/stretchr/testify/assert").
		AddImportedPackage(g.ModulePath + "/mocks").
		AddImportedPackage(g.ModulePath + "/service/api").
		AddImportedPackage("testing")
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
