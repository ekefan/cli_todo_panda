package store

import (
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
	jsonFile, err := fileExists()
	require.NotEmpty(t, jsonFile)
	require.NoError(t, err)
	//case two, file has been created, without tasks
	jsonFile, err = fileExists()
	require.NoError(t, err)
	require.Empty(t, jsonFile)
	defer jsonFile.Close()
	err = testStore.LoadTasks()
	require.NoError(t, err)
}

func TestSaveTask(t *testing.T) {
	err := testStore.SaveTasks()
	require.NoError(t, err)
}
