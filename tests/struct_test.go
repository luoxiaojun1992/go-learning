package tests

import (
	"encoding/json"
	"testing"
)

type DemoInterface interface {
	hello() string
}

type BaseStruct struct {
	Name string
	Age  int
}

type AnotherBaseStruct struct {
	AnotherName string
}

type DemoStruct struct {
	Foo string `json:"foo"`
	Bar string
	lowerStr string
	Number int
	string
	BaseStruct
	AnotherBase AnotherBaseStruct
	Age string
}

func NewDemoStruct(foo string) *DemoStruct {
	return &DemoStruct{Foo: foo}
}

func (base *BaseStruct) hello() string {
	return "Hello Base"
}

//func (base *BaseStruct) hello(a string) string {
//	return a
//}

//func (base *BaseStruct) Age(a string) string {
//	return a
//}

func (demo *DemoStruct) hello() string {
	return "Hello"
}

func GeneralHello(demo DemoInterface) string {
	return demo.hello()
}

func TestAccessField(t *testing.T) {
	demoStruct := new(DemoStruct)
	demoStruct.Name = "test struct"
	t.Log(demoStruct.Name)

	demoStruct2 := DemoStruct{}
	demoStruct2.Name = "test struct 2"
	t.Log(demoStruct2.Name)

	demoStruct3 := &DemoStruct{}
	demoStruct3.Name = "test struct 3"
	t.Log(demoStruct3.Name)

	demoStruct3.Foo = "test struct foo"
	t.Log(demoStruct3.Foo)

	demoStruct3.string = "test struct anonymous field"
	t.Log(demoStruct3.string)

	demoStruct3.AnotherBase.AnotherName = "test another base struct name"
	t.Log(demoStruct3.AnotherBase.AnotherName)

	demoStruct3.Name = "test base struct name"
	t.Log(demoStruct3.Name)

	demoStruct3.BaseStruct.Name = "test base struct name 2"
	t.Log(demoStruct3.BaseStruct.Name)

	demoStruct3.Age = "test struct age"
	t.Log(demoStruct3.Age)

	demoStruct3.BaseStruct.Age = 1
	t.Log(demoStruct3.BaseStruct.Age)

	demoStruct4 := DemoStruct{
		"test struct foo 2",
		"",
		"",
		0,
		"",
		BaseStruct{},
		AnotherBaseStruct{},
		"",
	}
	t.Log(demoStruct4.Foo)

	//error
	//t.Log(demoStruct4.BaseStruct.Bar)
}

func TestTag(t *testing.T) {
	jsonStr, errJsonEncoding := json.Marshal(&DemoStruct{
		Foo:         "",
		Bar:         "",
		lowerStr:    "",
		Number:      0,
		string:      "",
		BaseStruct:  BaseStruct{},
		AnotherBase: AnotherBaseStruct{},
		Age:         "",
	})
	if errJsonEncoding != nil {
		t.Log(errJsonEncoding)
		t.Fail()
	} else {
		t.Log(string(jsonStr))
	}
}

func TestStructMethod(t *testing.T) {
	demoStruct := new(DemoStruct)
	t.Log(demoStruct.hello())

	demoStruct2 := new(DemoStruct)
	t.Log(demoStruct2.BaseStruct.hello())
}

func TestStructInterface(t *testing.T) {
	t.Log(GeneralHello(new(DemoStruct)))
}
