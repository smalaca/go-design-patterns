package state

import "fmt"

type Recovery struct {

}

func (state *Recovery) changePassword(user *User, password string) {
	user.password = password
}

func (state *Recovery) changeLogin(user *User, login string) {
	fmt.Println("Recovery: changing login is forbidden")
}

func (state *Recovery) name() string {
	return "Recovery"
}