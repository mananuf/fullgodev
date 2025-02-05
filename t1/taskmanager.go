package main

import "fmt"

var TaskId uint = 0

type Status uint
type Tasks map[uint]*Task

var allTasks Tasks

type Task struct {
	id     uint
	title  string
	status Status
}

const (
	Pending Status = iota
	InProgress
	Done
)

func main() {
	addNewTask("bath")
	addNewTask("sleep")
	addNewTask("eat")
	updateTask(4, "status", Done)
	updateTask(2, "status", Done)
	updateTask(3, "status", InProgress)
	updateTask(2, "title", "i want to take a nap")
	fmt.Println(allTasks[2])
	fmt.Println(allTasks[3])
}

func addNewTask(title string) {
	if allTasks == nil {
		allTasks = make(Tasks)
	}
	TaskId = TaskId + 1

	allTasks[TaskId] = &Task{
		id:     TaskId,
		title:  title,
		status: Pending,
	}
}

func updateTask(taskId uint, key string, value interface{}) {
	if allTasks == nil {
		fmt.Println("No available Tasks at the moment")
		return
	}

	task, ok := allTasks[taskId]

	if !ok {
		fmt.Println("Task does not exits!")
		return
	}

	switch key {
	case "title":
		str, ok := value.(string)
		if !ok {
			return
		}
		task.title = str

		return
	case "status":
		status, ok := value.(Status)
		if !ok {
			return
		}
		task.status = status

		return
	}
}
