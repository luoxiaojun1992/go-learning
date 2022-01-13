package tests

import "testing"

func TestEmptyChannel(t *testing.T) {
	var ch1 chan int
	t.Log(ch1 == nil)

	ch2 := make(chan int)
	t.Log(ch2 == nil)
}

func TestChannelComparison(t *testing.T) {
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
