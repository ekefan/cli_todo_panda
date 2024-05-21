package task

import (
	"fmt"
	"os"
)

// task is the model for applications tasks
type task struct {
	Desc string
	Priority string
}

var Tasks []task

//CreateTask: creates a new task with desc, and priority
func CreateTask(desc, priority string) {
	newTask := task{
		Desc: desc,
		Priority: priority,

	}
	Tasks = append(Tasks, newTask)
}



//ListTasks: prints the list of tasks not completed to stdout
func ListTasks() {
	for i, tsk := range(Tasks) {
		fmt.Fprint(os.Stdout, i+1, tsk.Desc, tsk.Priority)
	}
}

func GetNext() {
	nextTask := Tasks[1]
	fmt.Fprint(os.Stdout, 2, nextTask.Desc, nextTask.Priority)
}


func DeleteTask(taskID int){
	Tasks = append(Tasks[:taskID], Tasks[taskID:]...)
	ListTasks()
}

func ClearAll() {
	Tasks = nil
	fmt.Fprintf(os.Stdout, "No task")
}