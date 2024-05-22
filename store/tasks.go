package store

import (
	"errors"
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
	correctArgLength(3, args, 1)

	if !checkPriority(args[2]){
		err := errors.New("<priority> must either be high/(H), low/(L) or none, (N)")
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


//ListTasks: prints the list of tasks not completed to stdout
func (s *Store)ListTasks(tasks []Task) {
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
		fmt.Fprintf(os.Stdout, "%d. %s %s\n",i+1, tsk.Desc, tsk.Priority)
	}
}


func (s *Store)DeleteTask(args []string, tasks []Task){ 
	correctArgLength(2, args, 2)
	taskID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	if len(tasks) == 0 {
		fmt.Fprintf(os.Stdout, "%s", "All tasks are completed\nAdd a tasks: add <desc> <priority>")
		return
	}

	idx := taskID - 1
	if taskID > len(s.Tasks) {
		fmt.Fprintf(os.Stdout, "No incomplete task with id, %d", taskID)
		return
	}
	task := s.Tasks[idx].Desc
	s.Tasks = nil
	s.Tasks = append(tasks[:idx], tasks[idx+1:]...)
	fmt.Fprintf(os.Stdout, "%s Completed\n%d Tasks Left:\n", task, len(s.Tasks))

	if !s.noTaskLeft() {
		s.ListTasks(s.Tasks)
	}
	err = s.SaveTasks()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store)ClearAll(tasks []Task) {
	if len(tasks) == 0 {
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
		
done: completes a tasks and removes it from completed tasks
  	usage: taskPanda done <taskID>

clear: removes every complete from tasks
  	usuage: clear all tasks from storage`
		fmt.Fprintf(os.Stdout, "%s", help)
}