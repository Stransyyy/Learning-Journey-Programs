package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Phone          string `json:"phone"`
	FavoritePhrase string `json:"custom1"`
}

func main() {
	data, derr := os.ReadFile("file.json")

	if derr != nil {
		log.Fatal(derr)
	}

	var p Person

	err := json.Unmarshal(data, &p)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Hi %s! You are %d years old. Your phone number is %s. Your favorite phrase is %s\n", p.Name, p.Age, p.Phone, p.FavoritePhrase)
}
