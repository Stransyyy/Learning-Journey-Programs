package main

import (
	//"flag"
	"fmt"
	//"math"
	//"math/rand"
)

func multi(x, y, c float64) float64 {

	var total float64
	total = x * y * c
	return total

}

func main() {

	//fmt.Printf("To log-in you must type your %+v/n")

	alan := User{"Alan", "alan@vitalitySouth.com", true, 18}
	fmt.Println(alan)
	fmt.Printf("\n Account info: %v \n Email: %v \n\n", alan.Name, alan.Email) //alan.Status)

	fmt.Println(multi(2, 3, 4))
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
