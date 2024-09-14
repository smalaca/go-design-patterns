package state

import "fmt"

type User struct {
	login string
	password string
	firstName string
	lastName string
}

func (user *User) changePassword(password string) {
	
}

func (user *User) changeLogin(login string) {
	
}

func (user *User) display() {
	fmt.Println("User:[" + 
		"Login: " + user.login + "; " +
		"password: " + user.password + "; " +
		"firstName: " + user.firstName + "; " +
		"lastName: " + user.lastName + "; " +
		"]")
}

func (user *User) block() {
	
}

func (user *User) recover() {
	
}

func (user *User) activate() {
	
}