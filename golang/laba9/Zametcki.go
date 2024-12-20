package laba9

type Zametcki interface {
	AddTask(description string)
	ShowTasks()
	UpdateTaskStatus(index int)
	DeleteTask(index int)
	SearchTask(word string)
}
