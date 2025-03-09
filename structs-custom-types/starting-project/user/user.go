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

// required to use the parameter in order to access the function; abc.User.OutputUser()
//type Admin struct {
//	userName string
//	password string
//	User     *User
//}

// Admin no required to user parameter; abc.OutputUser()
type Admin struct {
	userName string
	password string
	User
}

// NewUser constructor of a struct
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("field cannot be empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
	}, nil
}

func NewAdmin(username, password string, user *User) (*Admin, error) {
	if username == "" || password == "" {
		return nil, errors.New("field cannot be empty")
	}

	return &Admin{
		userName: username,
		password: password,
		User:     *user,
	}, nil
}

func (u *User) OutputUser() {
	// the shorthand
	fmt.Println(u.firstName, u.lastName, u.birthDate)

	// the manual dereference
	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
}
