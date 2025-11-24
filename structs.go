package main

import (
	"fmt"
	"structs/first-app/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	//***If a variable is a pointer to a struct, it automatically gets access to all methods that have a pointer receiver.***
	var appUser *user.User

	appUser, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println("Error creating user: ", err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	//pass the address of appUser to outputUserDetails instead of a copy of appUser
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
