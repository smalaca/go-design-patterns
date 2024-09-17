package state

type Active struct {

}

func (state *Active) changePassword(user *User, password string) {
	user.password = password
}

func (state *Active) changeLogin(user *User, login string) {
	user.login = login
}

func (state *Active) name() string {
	return "Active"
}