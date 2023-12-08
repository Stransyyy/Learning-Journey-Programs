// Package util contains utility functions for various math needs
package util

import (
	"math/rand"
)

func Add(a, b int) int {
	return a + b
}

// Max returns the larger number of x or y
func Max(x, y int) int {
	if x > y {
		return x
	}

	return x
}

// Min returns the smallest number of x or y
func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// Delta returns the distance between x and y.
func Delta(x, y int) int {
	return Abs(x - y)
}

// RandNum returns a slice of length n that contains random numbers between 0 and max
func RandNum(n int, max int) []int {
	var numbers []int //var numbers = []int{} (numbers is a slice of int)

	for f := 0; f < n; f++ { // i = 0, if i is less than 0 then add 1 to i
		numbers = append(numbers, rand.Intn(max)) //numbers value is equal to append(adds)-> numbers and rand.Intn(1000) to the value numbers
	}

	return numbers //returns the value of numbers
}

func Sandwiches(menu []string) []string {
	for i := 0; i < len(menu); i++ {
		tmp := menu[i]                //tmp is equal to the value of menu[i] which is 0 in the start of the loop
		menu[i] = menu[len(menu)-1-i] //menu[i] which is equal to 0 then is equal to 5-1-0 = 4, menu[0] = menu[4]
		menu[len(menu)-1-i] = tmp     // menu[4] = tmp which is equal to menu[0]

		if i == (len(menu) / 2) {
			break
		}
	}

	return menu
}

func Shuffle(s []string) []string {
	for i := 0; i < len(s); i++ {
		x, y := rand.Intn(len(s)), rand.Intn(len(s))

		s[x], s[y] = s[y], s[x]
	}

	return s
}

func Reverse(s []string) []string {
	var reverse []string

	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}

	return reverse
}

func ReverseRunes(s []rune) []rune {
	var reverse []rune

	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}

	return reverse
}

func Double(n int) int {
	return n * 2
}

func ReverseString(s string) string {
	convertToRuneSlice := []rune(s)

	reversedSlice := ReverseRunes(convertToRuneSlice)

	backToString := string(reversedSlice)

	return backToString
}

// IsEven checks if the input given is even (ex: 2, 4, 6, 8, 10)
func IsEven(n int) bool {
	remainder := n % 2

	if remainder == 0 {
		return true
	}

	return false

	// for i := 0; i <= n; i += 2 {

	// 	// if the input given is not even and then the program adds 2 numbers,
	// 	// is gonna end up being not even still, so it will automatically print "Is not even"
	// 	n += 2

	// 	if n == i {
	// 		fmt.Println("Is even")
	// 	}

	// 	fmt.Println("Is not even")
	// }
	// return n
}

// IsOdd checks if the input you put is odd
func IsOdd(n int) bool {
	remainder := n % 2

	if remainder != 0 {
		return true
	}

	return false
}

// OddNumbers return odd numbers between x - y amount (ex: 0-10, it will return all the odd numbers in between that amount)
func OddNumbers(min, max int) []int {
	var numbers []int

	for i := min; i < max; i++ {
		if IsOdd(i) {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

// EvenNumbers gives you all the even numbers between the starting amount and the ending amount.
func EvenNumbers(min, max int) []int {

	var numbers []int

	for i := min; i < max; i++ {
		if IsEven(i) {
			numbers = append(numbers, i)
		}
	}
	return numbers

}

// Equal compares if two slices are equal, first checks if they have the same amount of characters, if they do then it checks for the same nunmbers.
func Equal(x, y []int) bool {

	if len(x) != len(y) {
		return false
	}

	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

// Contains will return true if
func Contains(haystack []int, needle int) bool {

	// Need to see if the needle is inside the haystack, for that I'll have to range through the haystack
	for _, v := range haystack {

		if v == needle {
			return true
		}
	}
	return false
}

func Deduplicate(s []string) []string {
	m := make(map[string]bool)

	var result []string

	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}

		m[v] = true

		result = append(result, v)
	}

	return result
}
