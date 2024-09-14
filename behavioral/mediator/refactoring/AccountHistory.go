package mediator

import (
	"fmt"
	"strconv"
)

type AccountHistory struct {

}

func (history *AccountHistory) openHistoryFor(account Account) {
	fmt.Println("Account History: open history for account: " + strconv.Itoa(account.accountId))
	// some functionality
}

func (history *AccountHistory) verifyAccount(account Account) bool {
	fmt.Println("Account History: verify account: " + strconv.Itoa(account.accountId))
	// some functionality
	return true;
}

func (history *AccountHistory) verifyAccountRequest(accountRequest AccountRequest) bool {
	fmt.Println("Account History: verify account request: " + strconv.Itoa(accountRequest.accountRequestId))
	// some functionality
	return true;
}