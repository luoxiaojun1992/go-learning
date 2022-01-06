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
}
