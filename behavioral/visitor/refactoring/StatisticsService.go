package visitor

import "fmt"

type StatisticsService struct {
}

func (visitor *StatisticsService) registerTask(task *Task) {
	fmt.Println("Statistics about task: " + task.name)
}

func (visitor *StatisticsService) registerStory(story *Story) {
	fmt.Println("Statistics about story: " + story.name)
}

func (visitor *StatisticsService) registerProject(project *Project) {
	fmt.Println("Statistics about project: " + project.name)
}
