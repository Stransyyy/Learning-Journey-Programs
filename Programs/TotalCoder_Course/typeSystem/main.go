package main

import "fmt"

func main() {

	balances := map[string]int{
		"alice":   100,
		"bob":     200,
		"charlie": 300,
		"alan":    400,
	}

	for i, name := range balances {
		fmt.Printf("this is the index of the slice: %d and the name is %s \n", i, name)
	}
}
