// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {

	//access struct fields using the pointer dereference operator
	fmt.Println("first name: ", u.firstName)
	fmt.Println("last name: ", u.lastName)
	fmt.Println("birth date: ", u.birthDate)
	fmt.Println("created at: ", u.createdAt)

	//alternative way to access struct fields using pointer dereference, the above is the shortcut
	// fmt.Println((*u).firstName)

}

func main() {

	firstName := getUserData("Please enter your first name")
	lastName := getUserData("Please enter your last name")
	birthDate := getUserData("Please enter your birthDate (MM/DD/YYYY)")

	var appUser user

	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	//pass the address of appUser to outputUserDetails instead of a copy of appUser
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

/* // function now accepts a pointer to a user struct
func outputUserDetails(u *user) {

	//access struct fields using the pointer dereference operator
	fmt.Println("first name: ", u.firstName)
	fmt.Println("last name: ", u.lastName)
	fmt.Println("birth date: ", u.birthDate)
	fmt.Println("created at: ", u.createdAt)

	//alternative way to access struct fields using pointer dereference, the above is the shortcut
	// fmt.Println((*u).firstName)

}
*/

func getUserData(promptText string) string {

	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value

}
