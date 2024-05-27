package store

import (
	"slices"
	"strings"
	"errors"
	"fmt"
	"os"
	"bytes"
	"io"
)


// noTaskLeft: returns false if at least one task is incomplete
func (s *Store) noTaskLeft()bool{
	return len(s.Tasks) < 1
}

// correctArgLength: checks the lenth of args for each flag case
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

// checkPriority: checks if priority is included in pandas available priorities
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


func fileExists() (*os.File, error){
	var (
		newFile *os.File
		err error
	)
	_, errFound := os.Stat(filePath);
	if errors.Is(errFound, os.ErrNotExist) {
		newFile, err = os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil, err
		}
		if _, err := newFile.Write([]byte("null\n")); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil, err
		}
	}
	return newFile, nil
}

//stdRW defines the struct that defines the fields from redirecting os.Stdout to a writer/reader and restoring it back.
type stdRW struct {
	old *os.File
	reader *os.File
	writer *os.File
	chl chan string
}



func redirectStdOut(std *stdRW) error{
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	std.old = os.Stdout
	std.reader = r
	std.writer = w
	std.chl = make(chan string)

	os.Stdout = std.writer
	go func(reader *os.File) {
		var buffer bytes.Buffer
		io.Copy(&buffer, std.reader)
		std.chl <- buffer.String()
		close(std.chl)
	}(std.reader)
	return nil
}

func receiveFromStdOut(std *stdRW) string {
	std.writer.Close()
	os.Stdout = std.old
	stdOutput := <- std.chl
	std.reader.Close()
	return stdOutput
}
