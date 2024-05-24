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
		"fileName",
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
		"fileName",
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


}




func TestListTasks(t *testing.T){}
func TestDeleteTask(t *testing.T){}
func TestClearAll(t *testing.T){}
func TestHelp(t *testing.T){}