package mediator

import (
	"fmt"
	"strconv"
)

type AccountService struct {
	mediator Mediator
}

func (service *AccountService) createReuqest(requestId int, name string) {
	fmt.Println("Account Service: create request: " + strconv.Itoa(requestId) + "; " + name)

	accountRequest := AccountRequest{
		accountRequestId: requestId,
		name:             name,
	}

	if (service.mediator.isAcceptable(accountRequest)) {
		service.registerFromAccountRequest(accountRequest)
	}
}

func (service *AccountService) registerFromAccountRequest(accountRequest AccountRequest) {
	fmt.Println("Account Service: register from account request: " + strconv.Itoa(accountRequest.accountRequestId))
	// some functionality
}

func (service *AccountService) registerAccount(account Account) {
	fmt.Println("Account Service: register account: " + strconv.Itoa(account.accountId))
	// some functionality
}

func (service *AccountService) verify(account Account) bool {
	fmt.Println("Account Service: verify account: " + strconv.Itoa(account.accountId))
	// some functionality
	return true
}
