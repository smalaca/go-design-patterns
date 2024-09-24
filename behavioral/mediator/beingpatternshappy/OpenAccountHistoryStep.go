package mediator

type OpenAccountHistoryStep struct {
	next           ChainOfResponsibilityStep
	accountHistory *AccountHistory
}

func (step *OpenAccountHistoryStep) execute(account Account) {
	step.accountHistory.openHistoryFor(account)
	step.next.execute(account)
}
