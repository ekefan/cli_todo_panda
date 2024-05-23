package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func getFilePath(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath = fmt.Sprintf("%s/%s", home, name)
	return nil
}

//LoadTasks gets json data of the tasks and converts it the tasks array
func (s *Store) LoadTasks() error {
	err := getFilePath("task.json")
	if err != nil {
		//Print err to stdout
		return err
	}
	var jsonFile *os.File
	defer jsonFile.Close()
	jsonFile, errFound := fileExists()
	if errFound == nil {
		jsonFile, err = os.Open(filePath)
		if err != nil {
		//PrintError to stdout
		return err
	}

	}
	jsonTask, err := io.ReadAll(jsonFile)
	if err != nil {	
		//Print Error to stdOut
		return err
	}
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