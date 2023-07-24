package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func getComments(id int, responseCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{"One", "Two", "Three"}

	responseCh <- Response{
		data: comments,
		err:  nil,
	}

	wg.Done()
}

func getLikes(id int, responseCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	responseCh <- Response{
		data: 100,
		err:  nil,
	}

	wg.Done()
}

func getFriends(id int, responseCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	friends := []int{12, 13, 14}

	responseCh <- Response{
		data: friends,
		err:  nil,
	}

	wg.Done()
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		responseCh = make(chan Response, 3)
		wg         = &sync.WaitGroup{}
	)

	go getComments(id, responseCh, wg)
	go getLikes(id, responseCh, wg)
	go getFriends(id, responseCh, wg)

	// 3 processes using go routines
	wg.Add(3)
	// Block until the process count becomes 0
	wg.Wait()
	// Close the channel
	close(responseCh)

	userProfile := &UserProfile{}
	for res := range responseCh {
		if res.err != nil {
			return nil, res.err
		}

		switch msg := res.data.(type) {
		case int:
			userProfile.Likes = msg

		case []int:
			userProfile.Friends = msg

		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func main() {
	start := time.Now()
	user, err := handleGetUserProfile(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Info: %+v\n And took: %v\n", user, time.Since(start))
}
