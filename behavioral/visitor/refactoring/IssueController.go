package visitor

type IssueController struct {
	statisticsService *StatisticsService
	notificationService *NotificationService
}

func (controller *IssueController) processTask(task Task) {
	if (controller.notificationService != nil) {
		controller.notificationService.notifyAboutTask(&task)
	}

	if (controller.statisticsService != nil) {
		controller.statisticsService.registerTask(&task)
	}
}

func (controller *IssueController) processStory(story Story) {
	if (controller.notificationService != nil) {
		controller.notificationService.notifyAboutStory(&story)
	}

	if (controller.statisticsService != nil) {
		controller.statisticsService.registerStory(&story)
	}
}

func (controller *IssueController) processProject(project Project) {
	if (controller.statisticsService != nil) {
		controller.statisticsService.registerProject(&project)
	}
}

func (controller *IssueController) registerStatistics(statisticsService *StatisticsService) {
	controller.statisticsService = statisticsService
	controller.notificationService = nil
}

func (controller *IssueController) registerNotification(notificationService *NotificationService) {
	controller.notificationService = notificationService
	controller.statisticsService = nil
}
