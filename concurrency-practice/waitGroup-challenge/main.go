package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	wg.Add(1)
	go updateMessage("Update One")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Update Two")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Update Three")
	wg.Wait()
	printMessage()

}
