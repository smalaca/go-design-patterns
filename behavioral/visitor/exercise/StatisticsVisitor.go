package visitor

import "fmt"

type StatisticsVisitor struct {
}

func (visitor *StatisticsVisitor) visitTask(task *Task) {
	fmt.Println("Statistics about task: " + task.name)
}

func (visitor *StatisticsVisitor) visitStory(story *Story) {
	fmt.Println("Statistics about story: " + story.name)
}

func (visitor *StatisticsVisitor) visitProject(project *Project) {
	fmt.Println("Statistics about project: " + project.name)
}
