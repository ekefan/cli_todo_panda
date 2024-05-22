package store

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"log"
)

// Task is the model for applications tasks
type Task struct {
	Desc string
	Priority string
}

//CreateTask: creates a new Task with desc, and priority
func (s *Store)CreateTask(args []string) {
	if len(args) != 3 {
		err := errors.New("usage: taskPanda add <desc> <priority>")
		log.Fatal(err)
	}
	if !checkPriority(args[2]){
		err := errors.New("<priority> must be high (H), low (L) or none, (N)")
		log.Fatal(err)
	}
	newTask := Task{
		Desc: args[1],
		Priority: args[2],
	}

	s.Tasks = append(s.Tasks, newTask)
	err := s.SaveTasks()
	if err != nil {
		log.Fatal(err)
	}
}

func checkPriority(priority string) bool {
	priorities := []string{
		"none",
		"high",
		"low",
		"L",
		"H",
		"N",
	}
	return slices.Contains(priorities, priority)
}


//ListTasks: prints the list of tasks not completed to stdout
func (s *Store)ListTasks() {
	tasks := s.Tasks
	if len(tasks) == 0 {
		var err error
		tasks, err = s.LoadTasks()
		if err != nil {
		log.Fatal(err)
		}
		if len(tasks) == 0 {
			noTask := "No tasks added yet\n add: taskPanda add <desc> <priority>"
			fmt.Fprintf(os.Stdout, "%s", noTask)
			return
		}
	}
	for i, tsk := range(tasks) {
		fmt.Fprint(os.Stdout, i+1, tsk.Desc, tsk.Priority)
	}
}


func (s *Store)DeleteTask(taskID int){
	s.Tasks = append(s.Tasks[:taskID], s.Tasks[taskID:]...)
	s.ListTasks()
	err := s.SaveTasks()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store)ClearAll() {
	s.Tasks = nil
	s.SaveTasks()
	fmt.Fprintf(os.Stdout, "No Tasks added")
}

func (s *Store)Help(){
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
	On no arguments none priority tasks are returned
done:
  taskPanda done <taskID>
clear:
  clear all tasks from storage`
		fmt.Fprintf(os.Stdout, "%s", help)
}