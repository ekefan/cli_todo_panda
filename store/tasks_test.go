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
}
func TestClearAll(t *testing.T){}
func TestHelp(t *testing.T){}