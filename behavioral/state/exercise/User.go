package state

import "fmt"

type User struct {
	login string
	password string
	firstName string
	lastName string
	state UserState	
}

func (user *User) changePassword(password string) {
	user.state.changePassword(user, password)
}

func (user *User) changeLogin(login string) {
	user.state.changeLogin(user, login)
}

func (user *User) display() {
	fmt.Println("User:[" + 
		"Login: " + user.login + "; " +
		"password: " + user.password + "; " +
		"firstName: " + user.firstName + "; " +
		"lastName: " + user.lastName + "; " +
		"state: " + user.state.name() + "; " +
		"]")
}

func (user *User) block() {
	user.state = &Blocked{}
}

func (user *User) recover() {
	user.state = &Recovery{}
}

func (user *User) activate() {
	user.state = &Active{}
}