package main

import (
	"fmt"
	"os"
	"strconv"
)

var maxLevel int

func rec(id int) bool {
	if id > maxLevel {
		return true
	}
	done := make(chan bool)
	println(id)
	go func() {
		done <- rec(id + 1)
	}()
	<-done
	println(id, "done")
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s <INTEGER_MAX_LEVEL>\n", os.Args[0])
		os.Exit(1)
	}
	var err error
	if maxLevel, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s <INTEGER_MAX_LEVEL>\n", os.Args[0])
		os.Exit(2)
	}

	done := make(chan bool)
	go func() {
		done <- rec(1)
	}()
	println("main: waiting...")
	<-done
	println("main done")
}
