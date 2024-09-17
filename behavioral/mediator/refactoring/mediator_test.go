package mediator

import (
	"fmt"
	"testing"
)

func Test_Mediator(t *testing.T) {
	accountHistory := &AccountHistory{}
	accountService := &AccountService{accountHistory: *accountHistory}
	securityService := &SecurityService{
		accountHistory: *accountHistory,
		accountService: *accountService,
	}
	
	mediator := createAccountingMediator(accountHistory, accountService, securityService)
	controller := AccountController{
		mediator: mediator,
	}

	controller.accept(42);
	fmt.Println("---------------------")
	securityService.check(Account{accountId: 13})
	fmt.Println("---------------------")
	accountService.createReuqest(123, "yet another request")
}
