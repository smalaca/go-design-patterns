package state

type UserState interface {
	changePassword(user *User, password string)
	changeLogin(user *User, login string)
	name() string
}