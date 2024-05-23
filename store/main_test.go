package store

import (
	"log"
	"os"
	"testing"
)

var testStore *Store
func TestMain(m *testing.M){
	//create test store instance
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	testStore = NewStore()
	filePath = home + "/task_test.json"
	os.Exit(m.Run())
}