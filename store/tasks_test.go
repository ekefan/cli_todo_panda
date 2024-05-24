package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)
func TestCreateTask(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	newTask := randomTask()	
	//case one correct length of the field
	field := []string{
		"add",
		newTask.Desc,
		newTask.Priority,
	}
	ts.CreateTask(field)
	require.Equal(t, ts.Tasks[0].Desc, newTask.Desc)
	require.Equal(t, ts.Tasks[0].Priority, newTask.Priority)

	cs := NewStore()
	cs.LoadTasks()
	require.Equal(t, cs.Tasks[0].Desc, newTask.Desc)
	require.Equal(t, cs.Tasks[0].Priority, newTask.Priority)


	//case two incorrect length of the field
	nt := randomTask()
	field = []string{
		"add",
		nt.Desc,

	}
	var std = &stdRW{}
	err = redirectStdOut(std)
	require.NoError(t, err)
	ts.CreateTask(field)
	errString := receiveFromStdOut(std)
	require.NotEmpty(t, errString)
	errMsg := "need 3 args, usage: taskPanda add <desc> <priority>\n"
	require.Equal(t, errString, errMsg)

	//case three incorrect priority arg
	nt = randomTask()
	field = []string{
		"fileName",
		nt.Desc,
		"Wrong description",
	}
	var stdP = &stdRW{}
	err = redirectStdOut(stdP)
	require.NoError(t, err)
	ts.CreateTask(field)
	errString = receiveFromStdOut(stdP)
	require.NotEmpty(t, errString)
	errMsg = "<priority> must either be high/(H), low/(L) or none, (N)\n"
	require.Equal(t, errString, errMsg)

}




func TestListTasks(t *testing.T){
	//create store and keep tasks in them
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	nt := randomTask()	
	ts.Tasks = append(ts.Tasks, nt)
	ts.SaveTasks()
	field := []string{
		"fileName",
	}
	stdL := &stdRW{}
	err = redirectStdOut(stdL)
	require.NoError(t, err)
	ts.ListTasks(field, true)
	tasksString := receiveFromStdOut(stdL)
	require.NotEmpty(t, tasksString)

	// case two wrong field length
	field = []string{
		"fileName",
		"tasks",
	}
	stdL = &stdRW{}
	err = redirectStdOut(stdL)
	require.NoError(t, err)
	ts.ListTasks(field, true)
	tasksString = receiveFromStdOut(stdL)
	require.NotEmpty(t, tasksString)
	errMsg := "need 1 arg, usage: taskPanda tasks\n"
	require.Equal(t, tasksString, errMsg)
}
func TestDeleteTask(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	for i := 0; i < 4; i++ {
		nt := randomTask()	
		ts.Tasks = append(ts.Tasks, nt)
	}
	//case one correct id
	field := []string{
		"complete",
		"2",
	}
	stdD := &stdRW{}
	redirectStdOut(stdD)
	lenBefore := len(ts.Tasks)
	ts.DeleteTask(field)
	receiveFromStdOut(stdD)
	lenAfter := len(ts.Tasks)
	require.Equal(t, lenBefore - lenAfter, 1)
	//ID greater than highest task ID
	field = []string{
		"complete",
		"35",
	}
	stdD = &stdRW{}
	err = redirectStdOut(stdD)
	require.NoError(t, err)
	ts.DeleteTask(field)
	msgString := receiveFromStdOut(stdD)
	require.NotEmpty(t, msgString)
	require.Equal(t, msgString, "No incomplete task with id, 35")

	//Invalid Id
	field = []string{
		"complete",
		"-1",
	}
	stdD = &stdRW{}
	err = redirectStdOut(stdD)
	require.NoError(t, err)
	ts.DeleteTask(field)
	msgString = receiveFromStdOut(stdD)
	require.NotEmpty(t, msgString)
	require.Equal(t, msgString, "<taskID must contain only integers within 1 - 10")
}
func TestClearAll(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	//case one clear full storage with right length of argument
	for i := 0; i < 4; i++ {
		nt := randomTask()	
		ts.Tasks = append(ts.Tasks, nt)
	}
	ts.SaveTasks()
	stdC := &stdRW{}
	redirectStdOut(stdC)
	field := []string{
		"clear",
	}
	ts.ClearAll(field)
	msg := receiveFromStdOut(stdC)
	require.NotEmpty(t, msg)
	require.Equal(t, msg, "All Tasks cleared")

	// case two clear empty storage with right length of argum
	ts.Tasks = nil
	stdC = &stdRW{}
	redirectStdOut(stdC)
	ts.ClearAll(field)
	msg = receiveFromStdOut(stdC)
	require.NotEmpty(t, msg)
	require.Equal(t, msg, "No task added yet")

	// case three incorrect argument length
	stdC = &stdRW{}
	field = []string{
		"No correct",
		"argument length",
	}
	redirectStdOut(stdC)
	ts.ClearAll(field)
	msg = receiveFromStdOut(stdC)
	require.NotEmpty(t, msg)
	require.Equal(t, msg, "need 1 arg, usage: taskPanda clear\n")


}
func TestHelp(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	helpStr := `Usage taskPanda <tag> <args>
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
	stdH := &stdRW{}
	redirectStdOut(stdH)
	ts.Help()
	msg := receiveFromStdOut(stdH)
	require.NotEmpty(t, msg)
	require.Equal(t, msg, helpStr)			
}