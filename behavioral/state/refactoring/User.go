package state

import "fmt"

type User struct {
	login string
	password string
	firstName string
	lastName string
	state string	
}

func (user *User) changePassword(password string) {
	if (user.state == "Active" || user.state == "Recovery") {
		user.password = password
	} else {
		fmt.Println("Blocked: changing password is forbidden")
	}
}

func (user *User) changeLogin(login string) {
	if (user.state == "Active") {
		user.login = login
	} else {
		fmt.Println(user.state + ": changing password is forbidden")
	}
}

func (user *User) display() {
	fmt.Println("User:[" + 
		"Login: " + user.login + "; " +
		"password: " + user.password + "; " +
		"firstName: " + user.firstName + "; " +
		"lastName: " + user.lastName + "; " +
		"state: " + user.state + "; " +
		"]")
}

func (user *User) block() {
	user.state = "Blocked"
}

func (user *User) recover() {
	user.state = "Recovery"
}

func (user *User) activate() {
	user.state = "Active"
}