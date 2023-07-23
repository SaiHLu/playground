package main

import (
	"fmt"
	"time"
)

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Result : %d", n)
}

func sample() {

	msgCh := make(chan string, 128)

	msgCh <- "a"
	msgCh <- "b"
	msgCh <- "c"
	close(msgCh)

	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}
		fmt.Println("The message: ", msg)
	}

	// for msg := range msgCh {
	// 	fmt.Println("The message: ", msg)
	// }

	/**
	resultch := make(chan string) // unbuffered channel
	// resultch := make(chan string, 10) // buffered channel
	go func() {
		// resultch <- fetchResource(1)
		result := <-resultch
		fmt.Println("Result: ", result)
	}()

	resultch <- "hola"

	// fmt.Println("Hola")
	**/
}
