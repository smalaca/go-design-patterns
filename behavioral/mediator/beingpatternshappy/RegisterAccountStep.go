package mediator

type RegisterAccountStep struct {
	next           ChainOfResponsibilityStep
	accountService *AccountService
}

func (step *RegisterAccountStep) execute(account Account) {
	step.accountService.registerAccount(account)
	step.next.execute(account)
}
