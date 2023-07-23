package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
change timeout value here to test the cancellation of request
*/
const (
	timeoutMilli        = 100
	serverResponseMilli = 100
)

type Response struct {
	userId string
	err    error
}

func thirdPartyFetch() (string, error) {
	time.Sleep(time.Millisecond * serverResponseMilli)

	return "Response from server", nil
}

func fetchUser(ctx context.Context) (string, error) {
	// context, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeoutMilli)
	context, cancel := context.WithTimeout(ctx, time.Millisecond*timeoutMilli)
	defer cancel()

	fmt.Println("I am from context: ", ctx.Value("username"))

	resultCh := make(chan Response)

	go func() {
		res, err := thirdPartyFetch()
		resultCh <- Response{
			userId: res,
			err:    err,
		}
	}()

	select {
	/**
	Whether the timeout is exceeded or manually canceled (cancel())
	**/
	case <-context.Done():
		return "", context.Err()

	case res := <-resultCh:
		return res.userId, res.err
	}
}

func main() {
	start := time.Now()
	/*
		Passing value in context
	*/
	context := context.WithValue(context.Background(), "username", "Superman")
	result, err := fetchUser(context)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The response take %v and result %+v\n", time.Since(start), result)
}
