package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"one",
		"two",
		"three",
		"four",
	}
	wg.Add(len(words))

	for index, value := range words {
		go printSomething(fmt.Sprintf("Index %d => Value %s", index, value), &wg)
	}

	wg.Wait()

	fmt.Println("Final")
}
