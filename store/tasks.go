package store

import (
	"fmt"
	"os"
	"strconv"

	"log"
)

// Task is the model for applications tasks
type Task struct {
	Desc string
	Priority string
}

//CreateTask: creates a new Task with desc, and priority
func (s *Store)CreateTask(args []string) {
	if err := correctArgLength(3, args, 1); err != nil {
		return
	}


	if !checkPriority(args[2]){
		errMsg := "<priority> must either be high/(H), low/(L) or none, (N)"
		fmt.Fprintf(os.Stdout, "%s\n", errMsg)
		return
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


//ListTasks: prints the list of tasks not completed to stdout
func (s *Store)ListTasks(args []string, fromUser bool) { 
	if fromUser{
		if err := correctArgLength(1, args, 2); err != nil {
			return
		}
	} // flag 2 for listing task
	if len(s.Tasks) == 0 {
		noTask := "No tasks added yet\n add: taskPanda add <desc> <priority>"
		fmt.Fprintf(os.Stdout, "%s", noTask)
		return
	}

	for i, tsk := range(s.Tasks) {
		fmt.Fprintf(os.Stdout, "%d. %s %s\n",i+1, tsk.Desc, tsk.Priority)
	}
}


func (s *Store)DeleteTask(args []string){ 
	if err := correctArgLength(2, args, 3); err != nil {
		return
	}
	taskID, err := strconv.Atoi(args[1])
	if err != nil || taskID < 1{
		fmt.Fprintf(os.Stdout, "<taskID must contain only integers within 1 - 10")
		return
	}
	if len(s.Tasks) == 0 {
		fmt.Fprintf(os.Stdout, "%s", "All tasks are completed\nAdd a tasks: add <desc> <priority>")
		return
	}

	idx := taskID - 1
	if taskID > len(s.Tasks) {
		fmt.Fprintf(os.Stdout, "No incomplete task with id, %d", taskID)
		return
	}
	taskToComplete := s.Tasks[idx].Desc
	tasks := s.Tasks
	s.Tasks = nil
	s.Tasks = append(tasks[:idx], tasks[idx+1:]...)
	fmt.Fprintf(os.Stdout, "%s Completed\n%d Tasks Left:\n", taskToComplete, len(s.Tasks))

	if !s.noTaskLeft() {
		s.ListTasks([]string{}, false)
	}
	err = s.SaveTasks()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store)ClearAll(args []string) {
	if err := correctArgLength(1, args, 4); err != nil {
		return
	}
	if len(s.Tasks) == 0 {
		fmt.Fprintf(os.Stdout, "%s", "No task added yet")
		return
	}
	s.Tasks = nil
	s.SaveTasks()
	fmt.Fprintf(os.Stdout, "All Tasks cleared")
}

func (s *Store)Help(){
	help := `Usage taskPanda <tag> <args>
tags:
add :  adds a new task to incomplete tasks
	usage: taskPanda add <description> <priority>
	priority can be [high(H), low(L),  none(N)] -- can ignore caps

tasks: lists all uncompleted tasks
	usage: taskPanda tasks <tag>
	<tag>:
		when no tag -- returns all tasks
		-h -- returns high priority
		-l -- returns low priority
		-n -- returns none priority

complete: completes a tasks and removes it from completed tasks
	usage: taskPanda done <taskID>

clear: removes every tasks from storage
	usuage: taskPanda clear`
		fmt.Fprintf(os.Stdout, "%s", help)
}