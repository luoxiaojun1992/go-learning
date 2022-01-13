package tests

import "testing"

func TestEmptyMap(t *testing.T) {
	var eMap map[int]int
	t.Log(eMap == nil)

	map2 := make(map[int]int)
	t.Log(map2 == nil)
}

func TestMapComparison(t *testing.T) {
	//error mismatched key value type
	//map1 := make(map[int]int)
	//map2 := make(map[int]string)
	//map3 := make(map[string]int)
	//
	//map1 == map2
	//map1 == map3
	//map2 == map3

	var map1 map[string]string
	map2 := make(map[string]string)
	map1 = map2
	map2["foo"] = "bar"
	t.Log(map1["foo"])
}
