package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// "os"
// "io"

type Store struct {
	Tasks []Task
}


// Creates a new Instance of a store object
func NewStore() *Store {
	var newTask []Task
	return &Store{
		Tasks: newTask,
	}
}
var filePath string

//LoadTasks gets json data of the tasks and converts it the tasks array
func (s *Store) LoadTasks() error {
	home, err := os.UserHomeDir()
	if err != nil {
		//Print err to stdout
		return err
	}
	filePath = home + "/task.json"

	jsonFile, err := os.Open(filePath)
	if err != nil {
		//PrintError to stdout
		return err
	}
	defer jsonFile.Close()

	jsonTask, err := io.ReadAll(jsonFile)
	if err != nil {
		//Print Error to stdOut
		return err
	}

	// Unmarshall the json and store it to readTasks	
	err = json.Unmarshal(jsonTask, &s.Tasks)
	if err != nil {
		//Print Error to stdOUT
		return err
	}
	return nil
}

// SaveTasks converts tasks to json and writes data to a json file in the home dir
func (s *Store) SaveTasks() error {
	jsonData, err := json.Marshal(s.Tasks)
	if err != nil {
		//PrintError to std
		return err
	}
	//save the json data to the storage location
	writeErr := os.WriteFile(filePath, jsonData, 0666)
	if writeErr != nil {
		//Print Error to std
		return err
	}
	return nil
}


func (s *Store) noTaskLeft()bool{
	return len(s.Tasks) < 1
}


func correctArgLength(l int, args []string, flag int) error {
	var (
		errMsg string
		err error
	)
	if l != len(args) {
		switch flag {
		case 1: //create a new task
			errMsg = "need 3 args, usage: taskPanda add <desc> <priority>"
			fmt.Fprintf(os.Stdout, "%s\n", errMsg)
			err = errors.New(errMsg)
			return err
		
		case 2:
			errMsg = "need 1 arg, usage: taskPanda tasks"
			err = errors.New(errMsg)
			fmt.Fprintf(os.Stdout, "%s\n", errMsg)
			return err
		case 3:
			errMsg = "need 2 args, usage: taskPanda complete and <taskID>"
			err = errors.New(errMsg)
			fmt.Fprintf(os.Stdout, "%s\n", errMsg)
			return err
		case 4:
			errMsg = "need 1 arg, usage: taskPanda clear"
			err = errors.New(errMsg)
			fmt.Fprintf(os.Stdout, "%s\n", errMsg)
			return err
	    default:
			errMsg = ""
		}
	}
	return nil
}

func checkPriority(priority string) bool {
	priorities := []string{
		"none",
		"high",
		"low",
		"l",
		"h",
		"n",
	}
	return slices.Contains(priorities, strings.ToLower(priority))
}