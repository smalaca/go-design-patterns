package state

import "fmt"

type Blocked struct {

}

func (state *Blocked) changePassword(user *User, password string) {
	fmt.Println("Blocked: changing password is forbidden")
}

func (state *Blocked) changeLogin(user *User, login string) {
	fmt.Println("Blocked: changing login is forbidden")
}

func (state *Blocked) name() string {
	return "Blocked"
}