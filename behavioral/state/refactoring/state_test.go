package state

import (
	"testing"
)

func Test_State(t *testing.T) {
	originLogin := "spiderman"
	originPassword := "W1thGr8PowerComesGr8Respons!b!l!ty"
	user := User{
		login: originLogin,
		password: originPassword,
		firstName: "Peter",
		lastName: "Parker",
		state: "Active",
	}
	
	newPassword := "maybe less responsibility?"
	newLogin := "Venom"

	user.changePassword(newPassword)
	user.changeLogin(newLogin)
	user.display()

	user.block()
	user.changePassword(originPassword)
	user.changeLogin(originLogin)
	user.display()

	user.recover()
	user.changePassword("something new")
	user.changeLogin("spider ham")
	user.display()

	user.activate()
	user.changePassword(originPassword)
	user.changeLogin(originLogin)
	user.display()
}
