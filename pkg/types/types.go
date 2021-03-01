package types

import (
	"fmt"
	"github.com/aeramu/gocto/pkg/buffer"
)

type Struct struct {
	Name      string
	Variables []Variable
}

func NewStruct(name string) *Struct {
	return &Struct{Name: name}
}

func (t Struct) Render(b *buffer.Buffer) {
	b.Println("")
	b.Println("type %s struct {", t.Name)
	for _, variable := range t.Variables {
		b.Print("\t")
		variable.Render(b)
		b.Println("")
	}
	b.Println("}")
}

func (t *Struct) AddVariable(variable Variable) *Struct {
	t.Variables = append(t.Variables, variable)
	return t
}

type Interface struct {
	Name    string
	Methods []Method
}

func NewInterface(name string) *Interface {
	return &Interface{Name: name}
}

func (t Interface) Render(b *buffer.Buffer) {
	b.Println("")
	b.Println("type %s interface {", t.Name)
	for _, method := range t.Methods {
		b.Print("\t")
		method.Render(b)
		b.Println("")
	}
	b.Println("}")
}

func (t *Interface) AddMethod(method Method) *Interface {
	t.Methods = append(t.Methods, method)
	return t
}

type Function struct {
	*Method
	Receiver   *Variable
	Statements []string
}

func NewFunction(name string) *Function {
	method := NewMethod(name)
	return &Function{Method: method}
}

func NewFunctionFromMethod(method *Method) *Function {
	return &Function{Method: method}
}

func (t Function) Render(b *buffer.Buffer) {
	b.Println("")
	b.Print("func ")
	if t.Receiver != nil {
		b.Print("(")
		t.Receiver.Render(b)
		b.Print(") ")
	}
	t.Method.Render(b)
	b.Println(" {")
	for _, statement := range t.Statements {
		b.Print("\t")
		b.Println(statement)
	}
	b.Println("}")
}

func (t *Function) WithReceiver(variable Variable) *Function {
	t.Receiver = &variable
	return t
}

func (t *Function) AddParam(variable Variable) *Function {
	t.Params = append(t.Params, variable)
	return t
}

func (t *Function) AddReturn(ret string) *Function {
	t.Returns = append(t.Returns, ret)
	return t
}

func (t *Function) AddStatement(statement string, values ...interface{}) *Function {
	t.Statements = append(t.Statements, fmt.Sprintf(statement, values...))
	return t
}

type Method struct {
	Name    string
	Params  []Variable
	Returns []string
}

func NewMethod(name string) *Method {
	return &Method{Name: name}
}

func (t Method) Render(b *buffer.Buffer) {
	b.Print(t.Name)
	t.renderParams(b)
	b.Print(" ")
	t.renderReturns(b)
}

func (t *Method) AddParam(variable Variable) *Method {
	t.Params = append(t.Params, variable)
	return t
}

func (t *Method) AddReturn(ret string) *Method {
	t.Returns = append(t.Returns, ret)
	return t
}

func (t *Method) renderParams(b *buffer.Buffer) {
	b.Print("(")
	for i, param := range t.Params {
		param.Render(b)
		if i != len(t.Params)-1 {
			b.Print(", ")
		}
	}
	b.Print(")")
}

func (t *Method) renderReturns(b *buffer.Buffer) {
	if len(t.Returns) == 0 {
		return
	} else if len(t.Returns) == 1 {
		b.Print(t.Returns[0])
	} else {
		b.Print("(")
		for i, ret := range t.Returns {
			b.Print(ret)
			if i != len(t.Returns)-1 {
				b.Print(", ")
			}
		}
		b.Print(")")
	}
}

type Variable struct {
	Name string
	Type string
}

func NewVariable(name string, typeName string) Variable {
	return Variable{
		Name: name,
		Type: typeName,
	}
}

func (t Variable) Render(b *buffer.Buffer) {
	b.Print("%s %s", t.Name, t.Type)
}

type Header struct {
	PackageName      string
	ImportedPackages []string
}

func NewHeader(name string) *Header {
	return &Header{PackageName: name}
}

func (t *Header) Render(b *buffer.Buffer) {
	b.Println("package %s", t.PackageName)
	if len(t.ImportedPackages) > 0 {
		b.Println("")
		b.Println("import (")
		for _, name := range t.ImportedPackages {
			b.Println("\t\"%s\"", name)
		}
		b.Println(")")
	}
}

func (t *Header) AddImportedPackage(packageName string) *Header {
	t.ImportedPackages = append(t.ImportedPackages, packageName)
	return t
}
