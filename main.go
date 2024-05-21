package main


import (
	"fmt"
	"os"
)

func main() {
	osArgs := os.Args
	fmt.Println(osArgs)
	fmt.Println(len(osArgs))
}