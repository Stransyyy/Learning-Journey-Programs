package main

import (
	"fmt"
)

func eatCookies(messageCh chan string) {
	for message := range messageCh {
		fmt.Println("reading from channel", message)
	}
}

// A channel always blocks if its full
// A unbuffered channel is always full

func main() {
	messageCh := make(chan string, 128) // This is a unbuffered channel; how to make it buffered?
	//make(chan string, 1) <-- specify the size of the buffer
	//1. Write to a channel
	//2. Read to a channel

	// bufferedCh := make(chan string, 1) // buffered channel

	for i := 0; i < 100; i++ {
		messageCh <- fmt.Sprintf("hello %d\n", i)
	}

	close(messageCh)

	eatCookies(messageCh)

	// messageCh <- "Hello!" // blocking here

}
