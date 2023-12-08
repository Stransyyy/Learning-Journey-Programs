package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://vitalitysouth.com/"

func Request() {
	fmt.Println("VS web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type; %T \n", response)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Print(content)

	return
}

func test() {
	Request()
}
