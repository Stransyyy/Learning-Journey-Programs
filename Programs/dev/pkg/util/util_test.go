package util

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {

}

func TestAdd(t *testing.T) {
	type AddTest struct {
		Expect int
		Input1 int
		Input2 int
	}

	table := []AddTest{
		{
			Input1: 2,
			Input2: 3,
			Expect: 5,
		},
		{
			Input1: -2,
			Input2: 3,
			Expect: 1,
		},
		{
			Input1: 2,
			Input2: -3,
			Expect: -1,
		},
		{
			Input1: -2,
			Input2: -3,
			Expect: -5,
		},
		{
			Input1: 0,
			Input2: 0,
			Expect: 0,
		},
		{
			Input1: 1,
			Input2: 2,
			Expect: 3,
		},
	}

	for _, v := range table {
		got := Add(v.Input1, v.Input2)

		if got != v.Expect {
			t.Errorf("for inputs [%d, %d] expected %d; got %d", v.Input1, v.Input2, v.Expect, got)
		}
	}
}

func TestRandNum(t *testing.T) {
	n := 6
	max := 1000

	nums := RandNum(n, max)

	if len(nums) != n {
		t.Errorf("length of result is %d instead of %d", len(nums), n)
	}

	for _, v := range nums {
		if v < 0 || v > max {
			t.Errorf("result is outside valid range; got %d", v)
		}
	}
}

func TestIsEven(t *testing.T) {

	type IsEvenTest struct {
		Remainder int
		Even      bool
	}

	table := []IsEvenTest{
		{Remainder: 2, Even: true},
		{Remainder: 3, Even: false},
		{Remainder: 0, Even: true},
		{Remainder: -4, Even: true},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("IsEven %d should be %v", test.Remainder, test.Even), func(t *testing.T) {
			result := IsEven(test.Remainder)
			if result != test.Even {
				t.Errorf("IsEven(%d) returned %v, but expected %v", test.Remainder, result, test.Remainder)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	t.Parallel()

	type IsOddTest struct {
		Remainder int
		Odd       bool
	}

	table := []IsOddTest{
		{Remainder: 3, Odd: true},
		{Remainder: 2, Odd: false},
		{Remainder: 0, Odd: false},
		{Remainder: -3, Odd: true},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("IsEven %d should be %v", test.Remainder, test.Odd), func(t *testing.T) {
			result := IsOdd(test.Remainder)
			if result != test.Odd {
				t.Errorf("IsEven(%d) returned %v, but expected %v", test.Remainder, result, test.Remainder)
			}
		})
	}
}

func TestOddNumbers(t *testing.T) {

	type OddNumbersTest struct {
		Min    int
		Max    int
		OddNum []int
	}

	table := []OddNumbersTest{
		{Min: 0, Max: 10, OddNum: []int{1, 3, 5, 7, 9}},
		{Min: 0, Max: 20, OddNum: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}},
		{Min: 0, Max: 30, OddNum: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}},
		{Min: 0, Max: 40, OddNum: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("OddNumbers %d should be %v", test.Min, test.OddNum), func(t *testing.T) {
			result := OddNumbers(test.Min, test.Max)
			if result == nil {
				t.Errorf("OddNumbers(%d) returned %v, but expected %v", test.Min, result, test.OddNum)
			}
		})
	}

}

func TestEvenNumbers(t *testing.T) {

	type EvenNumbersTest struct {
		Min     int
		Max     int
		EvenNum []int
	}

	table := []EvenNumbersTest{
		{Min: 0, Max: 10, EvenNum: []int{2, 4, 6, 8, 10}},
		{Min: 0, Max: 20, EvenNum: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
		{Min: 0, Max: 30, EvenNum: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}},
		{Min: 0, Max: 40, EvenNum: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("EvenNumbers %d should be %v", test.Min, test.EvenNum), func(t *testing.T) {
			result := EvenNumbers(test.Min, test.Max)
			if result == nil {
				t.Errorf("EvenNumbers(%d) returned %v, but expected %v", test.Min, result, test.EvenNum)
			}
		})
	}

}

func TestEqual(t *testing.T) {

	type EqualTest struct {
		Slice1 []int
		Slice2 []int
		Equal  bool
	}

	table := []EqualTest{
		{Slice1: []int{1}, Slice2: []int{1}, Equal: true},
		{Slice1: []int{2}, Slice2: []int{2}, Equal: true},
		{Slice1: []int{3}, Slice2: []int{3}, Equal: true},
		{Slice1: []int{4}, Slice2: []int{4}, Equal: true},
		{Slice1: []int{1, 2, 3, 4, 5}, Slice2: []int{1, 2, 3, 4, 5}, Equal: true},
		{Slice1: []int{1, 2}, Slice2: []int{2, 4}, Equal: false},
		{Slice1: []int{1, 2, 3}, Slice2: []int{2, 4}, Equal: false},
		{Slice1: []int{1, 2, 3}, Slice2: []int{1, 3, 2}, Equal: false},
	}

	for _, test := range table {
		t.Run(fmt.Sprintf("Equal %d should be %v", test.Slice1, test.Equal), func(t *testing.T) {
			result := Equal(test.Slice1, test.Slice2)
			if result != test.Equal {
				t.Errorf("Test case failed: Input (%d, %v), got %v, expected %v", test.Slice1, result, test.Slice2, test.Equal)
			}
		})
	}

}

func TestContain(t *testing.T) {

	type ContainTest struct {
		haystack []int
		Needle   int
		Expected bool
	}
	farm := []ContainTest{
		{[]int{1, 2, 3, 4, 5}, 3, true},  // Test case 1: Needle is present in haystack.
		{[]int{1, 2, 3, 4, 5}, 6, false}, // Test case 2: Needle is not present in haystack.
		{[]int{}, 42, false},
		{[]int{1, 2, 3, 4, 5}, -3, false}, // Test case 3: Has an empty haystack, needle could not be found
	}

	for _, test := range farm {
		result := Contains(test.haystack, test.Needle)

		// Check if the result matches the expected value.
		if result != test.Expected {
			t.Errorf("Expected %v in %v to be %v, but got %v", test.Needle, test.haystack, test.Expected, result)
		}
	}
}
