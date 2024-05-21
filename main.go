package main


import (
	"fmt"
	"os"
)

func main() {
	osArgs := os.Args
	fields := osArgs[1:]
	fmt.Println(fields)
}