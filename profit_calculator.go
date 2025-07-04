package main

import (
	"fmt"
	"errors"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64
	
	
	// fmt.Println("Enter revenue: ")
	// fmt.Scan(&revenue)

	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Enter expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Enter expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Enter taxRate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Enter taxRate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("financials.txt", []byte(results), 0644)
}


func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Println(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		err := errors.New("Please enter a number greater than 0")
		return 0, err
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}