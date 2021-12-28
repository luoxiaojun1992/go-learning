package tests

import "testing"

func f1 () bool {
	//Error example
	//return (ret := true)

	return true
}

func TestFunc(t *testing.T) {
	t.Log(f1())
}
