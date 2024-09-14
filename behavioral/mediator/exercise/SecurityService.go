package mediator

import (
	"fmt"
	"strconv"
)

type SecurityService struct {
	mediator Mediator
}

func (service *SecurityService) check(account Account) bool {
	fmt.Println("Security Service: check account: " + strconv.Itoa(account.accountId))
	if (service.isValidAccount(account)) {
		return service.mediator.isValid(account);
	} else {
		return false;
	}
}

func (service *SecurityService) isValidAccount(account Account) bool {
	fmt.Println("Security Service: check if account is valid: " + strconv.Itoa(account.accountId))
	return true;
}

func (service *SecurityService) startMonitoring(account Account) {
	fmt.Println("Security Service: start monitoring: " + strconv.Itoa(account.accountId))
	// some functionality
}

