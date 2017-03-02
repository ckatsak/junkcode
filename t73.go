// Testing struct composition and ambiguous func selector
package main

import "fmt"

type Animal interface {
	Say() string
}

type Dog struct {
	sound string
}

func (d *Dog) Say() string {
	return d.sound
}

type Cat struct {
	sound string
}

func (c *Cat) Say() string {
	return c.sound
}

type Monster struct {
	*Dog
	*Cat
}

func main() {
	dog := &Dog{"woof"}
	fmt.Println(dog.Say())

	cat := &Cat{"meow"}
	fmt.Println(cat.Say())

	mon := &Monster{dog, cat}
	// ERROR because of ambiguity:
	// fmt.Println(mon.Say())
	// CORRECT:
	fmt.Printf("%s%s\n", mon.Dog.Say(), mon.Cat.Say())
}
