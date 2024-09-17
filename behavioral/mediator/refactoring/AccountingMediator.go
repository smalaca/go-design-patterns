package mediator

import "fmt"

type AccountingMediatior struct {
	accountHistory  *AccountHistory
	accountService  *AccountService
	securityService *SecurityService
}

func createAccountingMediator(accountHistory *AccountHistory, accountService *AccountService, securityService *SecurityService) Mediator {
	mediator := &AccountingMediatior{
		accountHistory:  accountHistory,
		accountService:  accountService,
		securityService: securityService,
	}

	return mediator
}

func (mediator *AccountingMediatior) newActiveAccount(account Account) {
	fmt.Println("MOVED to Mediator: Start")
	mediator.accountHistory.openHistoryFor(account)
	mediator.accountService.registerAccount(account)
	mediator.securityService.startMonitoring(account)
	fmt.Println("MOVED to Mediator: End")
}
