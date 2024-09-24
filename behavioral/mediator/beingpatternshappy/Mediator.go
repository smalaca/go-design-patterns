package mediator

type Mediator interface {
	isAcceptable(accountRequest AccountRequest) bool
	isValid(account Account) bool
	newActiveAccount(account Account)
}
