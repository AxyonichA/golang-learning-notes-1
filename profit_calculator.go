package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64
	
	
	// fmt.Println("Enter revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("Enter revenue: ")

	// fmt.Println("Enter expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInput("Enter expenses: ")

	// fmt.Println("Enter taxRate: ")
	// fmt.Scan(&taxRate)
	taxRate := getUserInput("Enter taxRate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Println(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}