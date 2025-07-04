package main

import (
	"fmt"
	"example.com/bank/fileops" //импортируем пакет, функции должны быть с большой буквы
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		// panic("Can't continue. Exiting...")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())
	// for i := 0; i < 2; i++ 
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile) 
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Print("Thanks for choosing our bank")
			return // stops the function, next code is unreachable
		}
	}
}
