package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/ekefan/cli_todo_panda/util"
	"github.com/stretchr/testify/require"
)


func randomTask() Task {
	return Task{
		Desc: util.RandomDesc(),
		Priority: util.RandomPriority(),
	}
}
func TestLoadTask(t *testing.T) {
	//case one empty file
	ts := NewStore()
	err := ts.SetFilePath("testTasks.json")//
	require.NoError(t, err)
	ts.LoadTasks()
	//first, check if no tasks exist in tasks
	require.Empty(t, ts.Tasks)

	//Case two file file with tasks
	newTask := randomTask()
	newTaskStr := fmt.Sprintf("[{\"Desc\":\"%s\", \"Priority\":\"%s\"}]\n", newTask.Desc, newTask.Priority)
	jsonFile, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	require.NoError(t, err)
	
	_, err = jsonFile.WriteString(newTaskStr)
	require.NoError(t, err)
	defer jsonFile.Close()
	ts.LoadTasks()
	require.NotEmpty(t, ts.Tasks)
	require.Equal(t, len(ts.Tasks), 1)
	require.Equal(t, ts.Tasks[0], newTask)
}

func TestSaveTask(t *testing.T) {
	ts := NewStore()
	ts.SetFilePath("saveTest.json")
	jsonFile, err := fileExists() //create newFile
	require.NoError(t, err)
	require.NotEmpty(t, jsonFile)
	newTask := randomTask()
	ts.Tasks = append(ts.Tasks, newTask)
	err = ts.SaveTasks() //save the newTask
	require.NoError(t, err)
	cs := NewStore()
	err = cs.LoadTasks()
	require.NoError(t, err)
	require.NotEmpty(t, cs.Tasks)
	require.Equal(t, ts.Tasks[0].Desc, cs.Tasks[0].Desc)
	require.Equal(t, ts.Tasks[0].Priority, cs.Tasks[0].Priority)
	jsonFile.Close()
}
