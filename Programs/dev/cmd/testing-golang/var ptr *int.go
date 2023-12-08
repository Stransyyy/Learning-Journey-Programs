package main

import (
	"fmt"
	"strings"
)

func mate(x, y *int) (z int) {
	z = *x + *y

	return z
}

func medio() {

	defer fmt.Println("Hola")

	suma := mate(new(int), new(int))

	suma = 9 + 5

	fmt.Println(suma)
}

func javier(a int) int {

	carlos := func(x int) int {
		return x + 1
	}(a)

	return a + carlos

}

type (
	Stack struct {
		top    *node
		lenght int
	}
	node struct {
		value interface{}
	}
)

func s() {
	var ab strings.Builder

	fmt.Println(ab)
	fmt.Println("Is this gonna work", ab.String())
	fmt.Println("Is this gonna work2", ab.Len())

}
