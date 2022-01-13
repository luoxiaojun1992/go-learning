package tests

import (
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestVar(t *testing.T) {
	var tv1 string
	t.Log(tv1)
	a, tv1 := "1", "2"
	t.Log(tv1)
	t.Log(a)
	tv1, a2, a, b := "1", "2", "3", "4"
	t.Log(a2)
	t.Log(b)

	//Error example
	//var tv2 string
	//t.Log(tv2)
	//tv2 := "1"
	//t.Log(tv2)

	tv3, err := "", errors.New("")
	t.Log(tv3)
	t.Log(err)

	tv4, err := "", errors.New("")
	t.Log(tv4)
	t.Log(err)

	tv5 := 0o12
	t.Log(tv5)

	tv6 := 012
	t.Log(tv6)

	tv7 := 0x12
	t.Log(tv7)

	tv8 := 0b10
	t.Log(tv8)

	t.Log(25 / 2)
	t.Log(float64(25) / 2)
	t.Log(float64(25) / float64(2))
	t.Log(25 / float64(2))

	tv9 := `a`
	t.Log(tv9)
	t.Log(reflect.TypeOf(tv9).String())

	var tv10, tv11 = 1, "11111"
	t.Log(tv10)
	t.Log(tv11)

	//Error example
	//var tv12 int, tv13 string = 1, "11111"
	var tv12, tv13 int = 1, 2
	t.Log(tv12)
	t.Log(tv13)

	t.Log([3]int{1,2,3} == [3]int{1,2,3})
	t.Log([...]int{1,2,3})
	t.Log([...]int{})
	t.Log([3]int{})
	t.Log([...][3]int{{1,2,3}, {1,2,3}})
	//Error example
	//t.Log([...][...]int{{1,2,3}, {1,2,3}})

	var i int
	t.Log(i == 0)
}
