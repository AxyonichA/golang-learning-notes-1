package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")
	var appUser *user.User
	// appUser = User{
	// 	// firstName: userFirstName,
	// 	// lastName: userLastName,
	// 	// birthDate: userBirthDate,
	// 	// createdAt: time.Now(),
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }
	appUser, err := user.New(userFirstName, userLastName, userBirthDate) // создаём при помощи конструктора
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "password")
	// вызывается метод User
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}