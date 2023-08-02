package main

import "testing"

func Test_mutex(t *testing.T) {
	msg = "One"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Three")
	wg.Wait()

	if msg != "Three" {
		t.Error("Incorrect value")
	}
}
