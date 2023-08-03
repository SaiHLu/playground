package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	now := time.Now()
	/*
		Since we have 5 philosophies, 5 * loop_counts(in this case 1000) = 5000 times
	*/
	for eatingCount := 1; eatingCount <= 1000; eatingCount++ {
		main()
	}

	if orderFinished != 5000 {
		t.Error("Expected 500 times")
	}

	t.Log("Finish In: ", time.Since(now).Microseconds())
}
