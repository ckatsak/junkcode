package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: $", os.Args[0], "<filename>")
		os.Exit(1)
	}

	cont, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: Reading the file")
		os.Exit(2)
	}

	fmt.Println(cont)

	return
}
