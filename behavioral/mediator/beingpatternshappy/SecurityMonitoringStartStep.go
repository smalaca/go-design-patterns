package mediator

type SecurityMonitoringStartStep struct {
	securityService *SecurityService
}

func (step *SecurityMonitoringStartStep) execute(account Account) {
	step.securityService.startMonitoring(account)
}