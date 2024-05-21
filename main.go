package main

import (
	"os"
	// "fmt"
	"strconv"
	"strings"

	"github.com/ekefan/cli_todo_panda/task"
)

const (
	add = "add"
	list = "list"
	del = "done"
	next = "next"
	clear = "clear"
	help = "help"

)
func main() {
	osArgs := os.Args
	fields := osArgs[1:]
	if len(fields) < 1 {
		task.Help()
		return
	}

	switch strings.ToLower(fields[0]) {
		case add: //takes 4 cli args
			task.CreateTask(fields[2], fields[3])
		case list: //takes 2 cli args
			task.ListTasks()
		case del: //takes 3 cli args
			taskID, _ := strconv.Atoi(fields[1])
			task.DeleteTask(taskID)
		case next:
			task.GetNext()
		case clear:
			task.ClearAll()
		case help:
			task.Help()
		default:
			task.Help()
	}
}