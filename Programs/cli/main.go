package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=e60f2ddf09bc4b3091c140905240705&q=Tupelo&aqi=no")
	if err != nil {
		fmt.Println("HTTP request failed: ", err)
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed: ", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}

	// file name for our json file
	fileName := "weather.json"

	// we check if it exist already
	if fi, err := os.Stat(fileName); err == nil {
		size := fi.Size()
		fmt.Printf("file already exists, size: %v bytes\n", size)
		if err := os.WriteFile(fileName, body, 0666); err != nil {
			fmt.Println("Failed to write to file: ", err)
			return
		}
		fmt.Println("File created successfully!")
	} else if err != nil {
		fmt.Println("Error checking for file")
	}

	// Pretty print API response to terminal
	// prettyBody, err := prettyJson(body)
	// if err != nil {
	// 	fmt.Println("Failed to pretty print JSON: ", err)
	// 	return
	// }
	// fmt.Printf("%s\n", prettyBody)

	// since we will make a cli, we won't need this pretty print

}

func prettyJson(body []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, body, "", "  ")
	return out.Bytes(), err
}
