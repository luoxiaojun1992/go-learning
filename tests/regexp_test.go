package tests

import (
	"regexp"
	"testing"
)

func TestFindAll(t *testing.T) {
	reg, _ := regexp.Compile("[a-z]*")
	t.Log(reg.FindAll([]byte("abc"), -1))
}
