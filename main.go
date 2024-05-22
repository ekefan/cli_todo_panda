package main

import (
	// "fmt"
	// "io"
	// "log"

	"log"
	"os"

	"strings"

	"github.com/ekefan/cli_todo_panda/store"
)

const (
	add = "add"
	list = "tasks"
	del = "done"
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
	//Load the tasks from the storage
	tasks, err := s.LoadTasks()
	if err != nil {
		log.Fatal(err)
	}
	switch strings.ToLower(fields[0]) {
		case add: //takes 4 cli args
			s.CreateTask(fields)
			return
		case list: //takes 2 cli args
			s.ListTasks(tasks)
			return
		case del: //takes 3 cli args
			s.DeleteTask(fields, tasks)
			return
		case clear:
			s.ClearAll(tasks)
			return
		case help:
			s.Help()
			return
		default:
			s.Help()
			return
	}
}