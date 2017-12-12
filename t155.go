package main

import "fmt"

func add4n5(s *[]int) {
	for i := 4; i <= 5; i++ {
		newS := append(*s, i)
		*s = newS
	}
	fmt.Println("func:", *s)
}

func wrong4n5(s *[]int) {
	for i := 4; i <= 5; i++ {
		newS := append(*s, i)
		s = &newS
	}
	fmt.Println("func:", *s)
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	fmt.Println("wrong:")
	wrong4n5(&x)
	fmt.Println("main:", x)

	fmt.Println("correct:")
	add4n5(&x)
	fmt.Println("main:", x)
}
