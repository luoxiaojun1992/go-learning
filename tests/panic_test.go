package tests

import "testing"

func TestRecover(t *testing.T) {
	defer func() {
		t.Log(recover())
	}()
	defer func() {
		t.Log(recover())
		panic("test panic 2")
	}()

	panic("test panic 1")
}
