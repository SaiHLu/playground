package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.setState(i + 1)
			wg.Done()
		}(i)
	}

	fmt.Println("Start")
	wg.Wait()
	fmt.Printf("state %+v\n", state)
	fmt.Println("Finish")
}
