package tests

import "testing"

type PointerTestDemoStruct struct {
	Name string
}

func TestStructPointer(t *testing.T) {
	var a *PointerTestDemoStruct
	a = &PointerTestDemoStruct{
		Name: "Pointer Test",
	}
	t.Log(a.Name)
}
