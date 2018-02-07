package main

import (
	"fmt"
)

type skata map[int]string

func newSkata() skata {
	return map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
}

func main() {
	k := newSkata()
	fmt.Printf("k[3] = %v\n", k[3])
}
