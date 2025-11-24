package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {

	//access struct fields using the pointer dereference operator
	fmt.Println("first name: ", u.firstName)
	fmt.Println("last name: ", u.lastName)
	fmt.Println("birth date: ", u.birthDate)
	fmt.Println("created at: ", u.createdAt)

	//alternative way to access struct fields using pointer dereference, the above is the shortcut
	// fmt.Println((*u).firstName)

}

func NewUser(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
	u.createdAt = time.Time{}
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
