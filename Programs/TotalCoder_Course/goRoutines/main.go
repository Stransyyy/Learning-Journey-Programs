package main

import (
	"fmt"
)

func eatCookie(messageCh chan string) {
	result := <-messageCh

	fmt.Println(result)
}

// A channel always blocks if its full
// A unbuffered channel is always full

func main() {
	messageCh := make(chan string) // This is a unbuffered channel; how to make it buffered?
	//make(chan string, 1) <-- specify the size of the buffer
	//1. Write to a channel
	//2. Read to a channel

	// bufferedCh := make(chan string, 1) // buffered channel

	go eatCookie(messageCh)

	messageCh <- "Hello!" // blocking here

}
