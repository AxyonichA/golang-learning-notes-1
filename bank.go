package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Balance file not found")
	}
	balanceText := string(data) // байты в строку
	balance, err :=strconv.ParseFloat(balanceText, 64) // строку в число с плавающей точкой
	
	if err != nil {
		return 1000, errors.New("Invalid balance format")
	}

	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		// panic("Can't continue. Exiting...")
	}

	fmt.Println("Welcome to Go Bank!")
	// for i := 0; i < 2; i++ 
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		
		var choise int
		fmt.Print("Your choise: ")
		fmt.Scan(&choise)
	
		/*
		if choise == 1 {

		} else if choise == 2 {

		}
		*/
		switch choise {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
	
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // stops current iteration and start next
			}
	
			accountBalance += depositAmount
	
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance) 
		case 3:
			var withdrawAmount float64
	
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // stop current iteration and start next
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than your balance.")
				continue
			}
	
			accountBalance -= withdrawAmount
	
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Print("Thanks for choosing our bank")
			return // stops the function, next code is unreachable
		}
	}

}
