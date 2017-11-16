package main

import "fmt"

type skata struct {
	x int
}

func newSkata() *skata {
	return &skata{}
}

func (s *skata) getX() int {
	s.x++
	return s.x
}

func main() {
	s := newSkata()
	for i := 0; i < 10; i++ {
		fmt.Println(s.getX())
	}
	fmt.Printf("%#v\n", s)
}
