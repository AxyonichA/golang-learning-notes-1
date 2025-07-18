package user

import (
	"fmt"
	"time"
	"errors"
)
type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}
// outputUserDetails теперь метод User
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}
	return &User{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
		}, nil
	}
func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:" Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}