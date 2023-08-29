package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func requestURL(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}

	return nil
}

/*
*
Using Goroutines
*
*/
func main() {
	startTime := time.Now()

	wg := sync.WaitGroup{}
	// requestStatusCh := make(chan string)

	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.similarweb.com/top-websites/",
		"https://mybloghosting.com/high-bandwidth-web-hosting-providers/",
		"https://www.webhostingsecretrevealed.net/blog/web-hosting-guides/how-much-bandwidth-does-your-site-really-need/",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			if err := requestURL(url); err != nil {
				log.Fatal(err)
				results[url] = "Failed"
			}

			results[url] = "Success"

			wg.Done()
		}(url)
	}

	wg.Wait()

	endTime := time.Since(startTime).Seconds()

	for url, result := range results {
		fmt.Println(url, result)
	}
	fmt.Println("Total Time: ", endTime)
}

/*
*
Without Goroutines
*
*/
/**
func main() {
	startTime := time.Now()

	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.similarweb.com/top-websites/",
		"https://mybloghosting.com/high-bandwidth-web-hosting-providers/",
		"https://www.webhostingsecretrevealed.net/blog/web-hosting-guides/how-much-bandwidth-does-your-site-really-need/",
	}

	for _, url := range urls {
		if err := requestURL(url); err != nil {
			log.Fatal(err)
			results[url] = "Failed"
		}

		results[url] = "Success"

	}

	endTime := time.Since(startTime).Seconds()

	for url, result := range results {
		fmt.Println(url, result)
	}
	fmt.Println("Total Time: ", endTime)
}
**/
