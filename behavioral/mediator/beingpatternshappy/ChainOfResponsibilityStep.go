package mediator

type ChainOfResponsibilityStep interface {
	execute(account Account)
}