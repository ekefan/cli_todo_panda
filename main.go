package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/ekefan/cli_todo_panda/store"
)

// tags available to Panda
const (
	add = "add"
	list = "tasks"
	del = "complete"
	clear = "clear"
	help = "help"
)
func main() {

	s := store.NewStore()
	osArgs := os.Args
	fields := osArgs[1:]
	if len(fields) < 1 {
		s.Help()
		return
	}
	// setFilePath
	srr := s.SetFilePath("task.json")
	if srr != nil {
		//Print err to stdout
		fmt.Fprintf(os.Stdout, "Couldn't set file path: %v", srr)
	}
	//Load the tasks from the storage
	err := s.LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Couldn't load Task from storage: %v",err)
		return
	}
	switch strings.ToLower(fields[0]) {
		case add: //takes 4 cli args
			s.CreateTask(fields)
			return
		case list: //takes 2 cli args
			s.ListTasks(fields, true)
			return
		case del: //takes 3 cli args
			s.DeleteTask(fields)
			return
		case clear:
			s.ClearAll(fields)
			return
		case help:
			s.Help()
			return
		default:
			s.Help()
			return
	}
}