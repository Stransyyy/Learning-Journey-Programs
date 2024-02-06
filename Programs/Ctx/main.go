package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.Background()
	UserID := 10
	val, err := fetchUserData(ctx, UserID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d", val)
	fmt.Printf("Time taken: %v", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, UserId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{val, err}

	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Fetching took too long: ", ctx.Err())
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100)

	return 666, nil
}
