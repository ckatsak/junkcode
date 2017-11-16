package main

import "fmt"

type Item interface {
	Equal(to Item) bool
	Less(than Item) bool
}

type Integer int

func (i Integer) Equal(j Item) bool {
	return i == j.(Integer)
}

func (i Integer) Less(j Item) bool {
	return i < j.(Integer)
}

func main() {
	s := []Integer{3, 2, 4, 5, 1}

	fmt.Println(s[0].Less(s[1]))
}

var _ Item = Integer(42)
