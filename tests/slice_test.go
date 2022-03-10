package tests

import "testing"

func TestSliceAssign(t *testing.T) {
	slice1 := []int{1,2,3}
	slice2 := []int{1,2,3,4}
	slice1 = slice2
	t.Log(slice1)
}

func TestArraySlice(t *testing.T) {
	arr1 := [...]int{1,2,3,4,5}
	slice1 := arr1[1:3:3]
	t.Log(len(arr1))
	t.Log(cap(arr1))
	t.Log(slice1)
	t.Log(len(slice1))
	t.Log(cap(slice1))
}

func TestMapSlice(t *testing.T)  {
	s := make([]map[string]string, 5)
	s[0] = map[string]string{"foo":"bar"}
	t.Log(s)
}

func TestEmptyArray(t *testing.T) {
	var arr1 [3][3]int
	t.Log(arr1)

	var arr2 [][3]int
	t.Log(arr2)
	t.Log(arr2 == nil)
}

func TestEmptySlice(t *testing.T) {
	var arr1 [][]int
	t.Log(arr1)
	t.Log(arr1 == nil)

	arr2 := make([][]int, 3)
	t.Log(arr2)
	t.Log(arr2 == nil)
	for _, el2 := range arr2 {
		t.Log(el2 == nil)
	}

	arr3 := make([]int, 3)
	t.Log(arr3)
	t.Log(arr3 == nil)

	arr4 := make([]int, 0)
	t.Log(arr4)
	t.Log(arr4 == nil)

	var arr5 []int
	t.Log(arr5 == nil)
	t.Log(append(arr5, 1))

	var arr6 []int
	t.Log(append([]int{1, 2, 3}, arr6...))

	var arr7 []int
	t.Log(len(arr7))

	var arr8 []int
	for i := range arr8 {
		t.Log(i)
	}
}
