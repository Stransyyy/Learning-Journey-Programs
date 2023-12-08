package main

import (
	"fmt"
)

func prueba(a int) int {

	for i := 0; i < 5; i++ {
		a = a + i //first chnage is gonna be 1 + 0 = 1
	}

	return a

}

func main() {

	b := prueba(1)

	fmt.Println(b)

	medio()

	a := javier(10)

	fmt.Println(a)

	s()

}

//Ctrl + n
