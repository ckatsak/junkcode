package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	maxLevel int
	step     int
)

func rec(id int) bool {
	if id > maxLevel {
		return true
	}
	done := make(chan bool)
	if id%step == 0 {
		println(id)
	}
	go func() {
		done <- rec(id + 1)
	}()
	<-done
	if id%step == 0 {
		println(id, "done")
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s <INTEGER_STEP> <INTEGER_MAX_LEVEL>\n", os.Args[0])
		os.Exit(1)
	}
	var err error
	if step, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s <INTEGER_STEP> <INTEGER_MAX_LEVEL>\n", os.Args[0])
		os.Exit(2)
	}
	if maxLevel, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ %s <INTEGER_STEP> <INTEGER_MAX_LEVEL>\n", os.Args[0])
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
