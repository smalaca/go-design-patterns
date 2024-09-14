package mediator

import "fmt"

type AccountController struct {
	mediator Mediator
}

func (controller *AccountController) accept(accountId int) {
	fmt.Println("Account Controller: accepting account")
	account := controller.findById(accountId);
	account.activate();

	controller.save(account);

	controller.mediator.newActiveAccount(account);
}

func (controller *AccountController) findById(accountId int) Account {
	return Account{
		accountId: accountId,
	}
}

func (controller *AccountController) save(account Account) {

}