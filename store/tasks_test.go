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
	//Create check store
	cs := NewStore()
	cs.LoadTasks()
	require.Equal(t, cs.Tasks[0].Desc, newTask.Desc)
	require.Equal(t, cs.Tasks[0].Priority, newTask.Priority)
	testCases := []struct{
		testName string
		field []string
		errMsg string
	}{
		{
			testName: "IncorrectFieldLength",
			field:	[]string{
				"add",
				randomTask().Desc,
			},
			errMsg: "need 3 args, usage: taskPanda add <desc> <priority>\n",
		},
		{
			testName: "WrongPriorityFormat",
			field:	[]string{
				"add",
				randomTask().Desc,
				"notPriority",
			},
			errMsg: "<priority> must either be high/(H), low/(L) or none, (N)\n",
		},
	}
	for _, tc := range(testCases) {
		t.Run(tc.testName, func(t *testing.T){
			var stdP = &stdRW{}
			err = redirectStdOut(stdP)
			require.NoError(t, err)
			ts.CreateTask(tc.field)
			errString := receiveFromStdOut(stdP)
			require.NotEmpty(t, errString)
			require.Equal(t, errString, tc.errMsg)
		})
	}

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
	testCases := []struct{
		name string
		field []string
		errMsg string
	}{
		{name: "Correct ID",field: []string{"complete","1",}},
		{
			name: "Id greater than number of task",
			field: []string{"complete","35",}, 
			errMsg:"No incomplete task with id, 35",
		},
		{
			name: "Invalid ID",
			field: []string{"complete","-1",}, 
			errMsg:"<taskID must contain only integers within 1 - 10",
		},
	}

	for _, tc := range(testCases) {
		t.Run(tc.name, func(t *testing.T){
			stdD := &stdRW{}
			if tc.name == "Correct ID" {
				redirectStdOut(stdD)
				lenBefore := len(ts.Tasks)
				ts.DeleteTask(tc.field)
				receiveFromStdOut(stdD)
				lenAfter := len(ts.Tasks)
				require.Equal(t, lenBefore - lenAfter, 1)	
			}else {
				err = redirectStdOut(stdD)
				require.NoError(t, err)
				ts.DeleteTask(tc.field)
				msgString := receiveFromStdOut(stdD)
				require.NotEmpty(t, msgString)
				require.Equal(t, msgString, tc.errMsg)
			}

		})
	}
}
func TestClearAll(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	for i := 0; i < 4; i++ {
		nt := randomTask()	
		ts.Tasks = append(ts.Tasks, nt)
	}
	testCase := []struct{
		name string
		field []string
		outputMsg string
	}{
		{name: "clear storage", field: []string{"clear"}, outputMsg: "All Tasks cleared"},
		{name: "clearing empty storage", field: []string{"clear"}, outputMsg: "No task added yet"},
		{name: "incorrect argument length", field: []string{"clear", "incorrect arg length"}, outputMsg: "need 1 arg, usage: taskPanda clear\n"},
	}
	for _, tc := range(testCase) {
		stdC := &stdRW{}
		if tc.name == 	"clear storage" {ts.SaveTasks()}
		if tc.name == "clearing empty storage" {ts.Tasks = nil}
		redirectStdOut(stdC)
		ts.ClearAll(tc.field)
		msg := receiveFromStdOut(stdC)
		require.NotEmpty(t, msg)
		require.Equal(t, msg, tc.outputMsg)
	}

}
func TestHelp(t *testing.T){
	ts := NewStore()
	err := ts.SetFilePath("CRUD_Test.json")
	require.NoError(t, err)
	helpStr := HELP_STR
	stdH := &stdRW{}
	redirectStdOut(stdH)
	ts.Help()
	msg := receiveFromStdOut(stdH)
	require.NotEmpty(t, msg)
	require.Equal(t, msg, helpStr)			
}