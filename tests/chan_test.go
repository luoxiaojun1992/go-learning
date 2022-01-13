package tests

import "testing"

func TestMapComparison(t *testing.T) {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int)
	ch2 = ch1
	t.Log(ch1 == ch2)

	go func() {
		t.Log(<-ch2)
	}()
	ch1 <- 1
}
