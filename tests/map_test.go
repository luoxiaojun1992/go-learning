package tests

import "testing"

func TestEmptyMap(t *testing.T) {
	var eMap map[int]int
	t.Log(eMap == nil)

	map2 := make(map[int]int)
	t.Log(map2 == nil)
}
