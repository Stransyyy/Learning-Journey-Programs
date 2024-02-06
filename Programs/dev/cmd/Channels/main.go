package main

import (
	"fmt"
	"strings"
	"time"
)

func sendData(ch chan<- string) {
	ch <- "Hello, "
	time.Sleep(1 * time.Second)
	ch <- "World!"
	time.Sleep(3 * time.Second)
	ch <- "Goodbye World!"
	close(ch)
}

func main() {
	ch := make(chan string)

	go sendData(ch)

	// Receiving data from the channel
	for msg := range ch {
		fmt.Print(msg)
		fmt.Print("\n")
		fmt.Println(strings.Split(msg, ""))
	}
}
