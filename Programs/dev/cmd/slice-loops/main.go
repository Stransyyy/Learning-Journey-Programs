package main

import (
	"fmt"

	"vs.local/pkg/util"
)

func main() {

	nums := util.RandNum(0, 10)

	fmt.Printf("[")
	for _, n := range nums { //Loop number 1
		fmt.Printf("%d ", n)
	}
	fmt.Printf("]\n")

	fmt.Println("This is a separation for loop number 1")

	fmt.Printf("[")
	for i := 0; i < len(nums); i++ { // Loop number 2
		fmt.Printf("%d ", nums[i])
	}
	fmt.Printf("]\n")

	fmt.Println("This is a separation for loop number 2")

	var reverse []int

	for i := len(nums) - 1; i >= 0; i-- {
		reverse = append(reverse, nums[i])
	}

	fmt.Println("Reversed:", reverse)

	menu := []string{"BLT", "Grilled Cheese", "Club", "Turkey", "Ham"}

	fmt.Println("Before:", menu)

	fmt.Println("Reversed:", util.Reverse(menu))
	fmt.Println("Shuffled:", util.Shuffle(menu))
	fmt.Println("Shuffled again:", util.Shuffle(menu))

	tmp := util.ReverseString("hello")

	fmt.Println(tmp)
}

///home/alan/src/dev/cmd/slice-loops/main.go
