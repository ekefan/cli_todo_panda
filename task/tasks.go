package task

import "fmt"

// task is the model for applications tasks
type task struct {
	Desc string
	Priority string
}

var Tasks []task

func CreateTask(desc, priority string, completed bool) {
	newTask := task{
		Desc: desc,
		Priority: priority,

	}
	Tasks = append(Tasks, newTask)
}


func ListTasks() {
	for i, tsk := range(Tasks) {
		fmt.Println(i+1, tsk.Desc, tsk.Priority)
	}
}

func GetNext() {
	nextTask := Tasks[1]
	fmt.Println(2, nextTask.Desc, nextTask.Priority)
}

func DeleteTask(taskID int){
	Tasks = append(Tasks[:taskID], Tasks[taskID:]...)
	ListTasks()
}