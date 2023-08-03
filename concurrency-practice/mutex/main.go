package main

import (
	"fmt"
	"sync"
)

// var (
// 	msg string
// 	wg  sync.WaitGroup
// )

// func updateMessage(s string) {
// 	defer wg.Done()

// 	msg = s
// }

// func main() {
// 	wg.Add(2)
// 	go updateMessage("One")
// 	go updateMessage("Two")
// 	wg.Wait()

// 	fmt.Println(msg)
// }

var (
	msg string
	wg  sync.WaitGroup
	mu  sync.Mutex
)

func updateMessage(s string) {
	defer wg.Done()
	defer mu.Unlock()

	mu.Lock()
	msg = s
}

func main() {

	wg.Add(2)
	go updateMessage("One")
	go updateMessage("Two")
	wg.Wait()

	fmt.Println(msg)
}
