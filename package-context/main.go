package main

import (
	"fmt"
	"log"
	"time"
)

func thirdPartyFetch() (string, error) {
	time.Sleep(time.Millisecond * 200)

	return "Response from server", nil
}

func main() {
	start := time.Now()
	result, err := thirdPartyFetch()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The response take %v and result %+v\n", time.Since(start), result)
}
