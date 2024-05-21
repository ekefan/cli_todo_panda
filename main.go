package main

import (
	"os"
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

)
func main() {
	osArgs := os.Args
	fields := osArgs[1:]
	
	switch strings.ToLower(fields[1]) {
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
	}
}