package main

import "fmt"

type Foo struct {
	x, y int
	s    string
}

func (f *Foo) x42() {
	f.x = 42
}

func (f *Foo) y666() {
	f.y = 666
}

func (f *Foo) syo() {
	f.s = "yo"
}

func (f Foo) shi() {
	f.s = "hi"
}

func main() {
	foo := Foo{}
	fmt.Printf("%#v\n", foo)
	foo.x42()
	foo.y666()
	foo.syo()
	fmt.Printf("%#v\n", foo)

	foo.shi()
	fmt.Printf("%#v\n", foo)

	fmt.Println(`-----------------------------------------`)

	foo2 := &Foo{}
	fmt.Printf("%#v\n", foo2)
	foo2.x42()
	foo2.y666()
	foo2.syo()
	fmt.Printf("%#v\n", foo2)

	foo2.shi()
	fmt.Printf("%#v\n", foo2)
}
