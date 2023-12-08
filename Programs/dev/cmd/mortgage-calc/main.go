package main

import (
	"fmt"
	"math"
)

func main() {

	var apr float64        // annual percentage rate
	var loanAmount float64 // loan amount
	var loanTerm float64   // loan term in years

	fmt.Println("Enter the ammount you owe ($): ")
	_, err := fmt.Scan(&loanAmount)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Enter the annual percentage rate (%): ")
	_, err = fmt.Scan(&apr)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Enter the loan term in years: ")
	_, err = fmt.Scan(&loanTerm)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	n := math.Ceil(loanTerm * 12) // convert years to months
	r := apr / 12
	a := loanAmount * (r / 100) * math.Pow((1+r/100), n) / (math.Pow((1+r/100), n) - 1)
	// convert apr to monthly interest rate
	monthlyInterestRate := (apr / 100) / 12
	fmt.Println("Monthly mortgage rate: ", a)
	fmt.Printf("%.2f%% over %.0f years:\n", apr, loanTerm)
	fmt.Print("Monthly Interest Rate: ", monthlyInterestRate, "\n")

}
