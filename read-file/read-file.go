package main

import (
	"fmt"
	"io"
	"os"
)



func main() {
  // Get list of args
	fmt.Println(os.Args)
	file := os.Args[1]
	openedFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file!")
		os.Exit(1)
	}
	io.Copy(os.Stdout, openedFile)
}
