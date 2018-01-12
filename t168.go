package main

import "fmt"

func main() {
	test1()
	test2()
}

func test1() {
	fmt.Println("test1: Entering")
	defer fmt.Println("test2: Exiting")
	fmt.Println("test1: a")
	goto sameNameLabel
	fmt.Println("test1: b")
sameNameLabel:
	fmt.Println("test1: c")
}

func test2() {
	fmt.Println("test2: Entering")
	defer fmt.Println("test2: Exiting")
	fmt.Println("test2: a")
	goto sameNameLabel
	fmt.Println("test2: b")
sameNameLabel:
	fmt.Println("test2: c")
}
