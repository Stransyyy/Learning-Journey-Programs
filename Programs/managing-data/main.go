package main

import (
	"github.com/Stransyyy/Learning-Journey-Programs/Programs/managing-data/data"
)

func main() {
	s, err := data.Get()
	if err != nil {
		panic(err)
	}

	println(s)
}
