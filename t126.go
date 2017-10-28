package main

import "fmt"

type myType struct {
	k, v int
}

func (t *myType) yo() (int, int) {
	if t == nil {
		return -1, -1
	}
	return t.k, t.v
}

func main() {
	t1 := &myType{k: 42, v: 42}
	k, v := t1.yo()
	fmt.Printf("t1: (%d, %d)\n", k, v)

	var t2 *myType = nil
	k, v = t2.yo()
	fmt.Printf("t2: (%d, %d)\n", k, v)
}
