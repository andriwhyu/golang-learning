package main

import (
	"example.com/starting-project/user"
	"fmt"
)

type customType string

func (c customType) log() {
	fmt.Println(c)
}

func main() {
	firstName := getUserData("first name: ")
	lastName := getUserData("last name: ")
	birthDate := getUserData("birthday (DD/MM/YYYY): ")

	userData, err := user.NewUser(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	userData2, err := user.NewUser("Budi", "Doremi", "25/08/2000")
	if err != nil {
		fmt.Println(err)
		return
	}

	userData.OutputUser()

	fmt.Println(&userData, &userData2)
	fmt.Println(userData, userData2)

	userAdmin, err := user.NewAdmin("admin", "admin", userData)
	if err != nil {
		fmt.Println(err)
		return
	}

	//userAdmin.User.OutputUser()
	userAdmin.OutputUser()

	// using type for primitive data type.
	var valCustom customType

	valCustom = "Hi, there!"
	valCustom.log()
}

func getUserData(prompt string) string {
	var userData string

	fmt.Print(prompt)
	_, err := fmt.Scanln(&userData)
	if err != nil {
		return ""
	}

	return userData
}
