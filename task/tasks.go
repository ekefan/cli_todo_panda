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

func Help(){
	help := `Usage taskPanda <tag> <args>
tags: 
add : 
taskPanda add <description> <priority>
  priority can be [high, low,  none]
list:
  taskPanda list <args>
  args:
	all -- returns all tasks
	high -- returns high priority
	low -- returns low priority
	On no arguments none priority task are returned
done:
  taskPanda done <taskID>
clear:
  clear all tasks from storage`
		fmt.Fprintf(os.Stdout, help)
}