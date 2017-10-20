package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "main: ", log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func main() {
	fmt.Println("--------------------------\n")

	wdstr, err := os.Getwd()
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	// fmt.Printf("working directory: \"%s\".\n", wdstr)
	logger.Printf("working directory: %q.\n", wdstr)

	wd, err := os.Open(wdstr)
	if err != nil {
		logger.Println(err)
		os.Exit(2)
	}

	content, err := wd.Readdirnames(0)
	if err != nil {
		logger.Println(err)
		os.Exit(3)
	}

	fmt.Printf("%q content:\n", wdstr)
	for _, f := range content {
		fmt.Printf("\t%s\n", f)
	}

	fmt.Println("--------------------------\n")

	a, err := os.Open("a")
	if err != nil {
		logger.Println(err)
		os.Exit(2)
	}

	ac, err := a.Readdirnames(0)
	if err != nil {
		logger.Println(err)
		os.Exit(3)
	}

	fmt.Printf("%q content:\n", "a")
	for _, f := range ac {
		fmt.Printf("\t%s\n", f)
		if fstr, err := filepath.Abs(f); err != nil {
			logger.Println(err)
		} else {
			fmt.Println(fstr)
		}
	}

	fmt.Println("--------------------------\n")
}
