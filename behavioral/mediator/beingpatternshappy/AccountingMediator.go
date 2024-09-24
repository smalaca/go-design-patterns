package mediator

import (
	"fmt"
	"strconv"
)

type AccountingMediator struct {
	accountHistory  *AccountHistory
	accountService  *AccountService
	securityService *SecurityService
	step            ChainOfResponsibilityStep
}

func createAccountingMediator(accountHistory *AccountHistory, accountService *AccountService, securityService *SecurityService) AccountingMediator {
	// Facade
	// Command
	// Interpretera
	step := &OpenAccountHistoryStep{
		accountHistory: accountHistory,
		next: &RegisterAccountStep{
			accountService: accountService,
			next: &SecurityMonitoringStartStep{
				securityService: securityService,
			},
		},
	}

	mediator := AccountingMediator{
		accountHistory:  accountHistory,
		accountService:  accountService,
		securityService: securityService,
		step:            step,
	}

	securityService.mediator = &mediator
	accountService.mediator = &mediator

	return mediator
}

func (mediator *AccountingMediator) isAcceptable(accountRequest AccountRequest) bool {
	fmt.Println("Mediator: is acceptable account requst: " + strconv.Itoa(accountRequest.accountRequestId))
	return mediator.accountHistory.verifyAccountRequest(accountRequest)
}

func (mediator *AccountingMediator) isValid(account Account) bool {
	fmt.Println("Mediator: is valid account: " + strconv.Itoa(account.accountId))
	return mediator.accountHistory.verifyAccount(account) && mediator.accountService.verify(account)
}

func (mediator *AccountingMediator) newActiveAccount(account Account) {
	fmt.Println("Mediator: new active account: " + strconv.Itoa(account.accountId))
	mediator.step.execute(account)
}
