package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const url1 = "https://lco.dev"
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	response, err := http.Get(url1)
	checkErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkErr(err)
	content := string(databytes)

	fmt.Print(content)
}

func parsing(err error) {
	result, _ := url.Parse(myurl)
	checkErr(err)

	fmt.Print(result.Scheme)

}

func checkErr(err error) error {
	if err != nil {
		return err
	}
	return err
}
