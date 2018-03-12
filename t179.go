package main

import "fmt"

type MyType struct {
	x int
}

func (t *MyType) Method(y int) int {
	return t.x + y
}

func MyFunc(z int, f func(int) int) {
	fmt.Printf("%d\n", f(z))
}

func main() {
	t := &MyType{42}
	for i := 0; i < 10; i++ {
		MyFunc(i, t.Method)
	}
}
