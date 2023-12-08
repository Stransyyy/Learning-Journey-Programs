package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func areaTraingle(t Triangle) float64 {
	return t.base * t.height / 2
}

func main() {

	base1 := 10.00
	height := 5.00

	mytriangle := Triangle{base1, height}

	area := areaTraingle(mytriangle)

	fmt.Printf("\nThe values are %.2f as height, and %.2f as the base. This is the area of your triangle: %.2f", height, base1, area)
}
