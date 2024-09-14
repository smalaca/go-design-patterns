package mediator

import (
	"fmt"
	"strconv"
)

type AccountController struct {
	accountHistory  AccountHistory
	accountService  AccountService
	securityService SecurityService
}

func (controller *AccountController) accept(accountId int) {
	fmt.Println("Account Controller: accepting account")
	account := controller.findById(accountId);
	account.activate();

	controller.save(account);

	controller.newActiveAccount(account);
}

func (controller *AccountController) newActiveAccount(account Account) {
	fmt.Println("To be moved to Mediator: new active account: " + strconv.Itoa(account.accountId))
	controller.accountHistory.openHistoryFor(account)
	controller.accountService.registerAccount(account)
	controller.securityService.startMonitoring(account)
}

func (controller *AccountController) findById(accountId int) Account {
	return Account{
		accountId: accountId,
	}
}

func (controller *AccountController) save(account Account) {

}