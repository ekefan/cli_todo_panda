package store

import (
	"encoding/json"
	"errors"
	"io"
	"log"
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
func (s *Store) LoadTasks() ([]Task, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		//Print err to stdout
		return nil, err
	}
	filePath = home + "/task.json"

	jsonFile, err := os.Open(filePath)
	if err != nil {
		//PrintError to stdout
		return nil, err
	}
	defer jsonFile.Close()

	jsonTask, err := io.ReadAll(jsonFile)
	if err != nil {
		//Print Error to stdOut
		return nil, err
	}

	// Unmarshall the json and store it to readTasks	
	err = json.Unmarshal(jsonTask, &s.Tasks)
	if err != nil {
		//Print Error to stdOUT
		return nil, err
	}
	return s.Tasks, nil
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


func correctArgLength(l int, args []string, flag int) {
	var err error
	if l != len(args) {
		switch flag {
		case 1: //create a new task
			err= errors.New("usage: taskPanda add <desc> <priority>, need 4 args")
		
		case 2:
			err = errors.New("usage: done <taskID, need only 2 args")
		}
		log.Fatal(err)
	}
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