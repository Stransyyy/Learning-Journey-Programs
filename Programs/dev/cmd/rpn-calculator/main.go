package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stack []int

func (st *stack) Push(a int) { //pushes an element into the stack
	*st = append(*st, a)
}

func (st *stack) Pop() (int, error) { // pops(deletes) the last element from the stack
	if st.IsEmpty() {
		return 0, errors.New("stack underflow")
	}

	// Get the last element
	a := (*st)[len(*st)-1]

	// Remove the last element
	*st = (*st)[:len(*st)-1]

	return a, nil
}

func (st *stack) Peek() (int, error) { // checks the last element of the stack
	if st.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return (*st)[len(*st)-1], nil
}

func (st *stack) IsEmpty() bool { // checks if the stack is empty
	return len(*st) == 0
}

func (st *stack) Clear() { //clears the hwole stack
	*st = stack{}
}

func (st *stack) Count() int { // counts the number of elements in the stack
	return len(*st)
}

func reversePolishNotation(q string) (int, error) {
	tokens := strings.Fields(q)
	st := &stack{}

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			st.Push(num)
		} else {
			b, err := st.Pop()
			if err != nil {
				return 0, err
			}
			a, err := st.Pop()
			if err != nil {
				return 0, err
			}

			switch token {
			case "+":
				st.Push(a + b)
			case "-":
				st.Push(a - b)
			case "*":
				st.Push(a * b)
			case "/":
				if b == 0 { // checks if the divisor is zero
					return 0, errors.New("division by zero") // if ti is zero, it returns an error
				}
				st.Push(a / b)
			default:
				return 0, fmt.Errorf("invalid operation: %s", token)
			}
		}
	}

	result, err := st.Pop()
	if err != nil {
		return 0, err
	}

	if !st.IsEmpty() {
		return 0, errors.New("extra elements in the stack")
	}

	return result, nil
}

func main() {
	q := "3 4 +"
	result, err := reversePolishNotation(q)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
