package main

//"fmt"

type stack []int // this is the stack type

func (st *stack) push(a int) {
	*st = append(*st, a)
}

// pop removes and returns the top element of the stack except when the stack is empty, in which case it returns 0
func (st *stack) pop() (int, bool) {

	if st.IsEmpty() {
		return 0, false
	}

	// get  final element
	a := (*st)[len(*st)-1]

	// remove final element
	*st = (*st)[:len(*st)-1]

	return a, true
}

func (st *stack) IsEmpty() bool {
	return len(*st) == 0
}

func clear(st *stack) {
	for !st.IsEmpty() {
		st.pop()
	}
}

func (st *stack) peek(int, bool) {
	if st.IsEmpty() {
	}

	return
}
