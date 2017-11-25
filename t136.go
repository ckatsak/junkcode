package main

import "fmt"

func main() {
	s1 := make([]int, 0)
	s2 := make([]int, 50)
	s3 := make([]int, 100)
	s4 := make([]int, 500)
	s5 := make([]int, 100)
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d\n", len(s2), cap(s2))
	fmt.Printf("len(s3): %d, cap(s3): %d\n", len(s3), cap(s3))
	fmt.Printf("len(s4): %d, cap(s4): %d\n", len(s4), cap(s4))
	fmt.Printf("len(s5): %d, cap(s5): %d\n", len(s5), cap(s5))
	for i := 0; i < 100; i++ {
		s1 = append(s1, i)
		s2 = append(s2, i)
		s3 = append(s3, i)
		s4 = append(s4, i)
		s5[i] = i
	}
	fmt.Println("--------------------------")
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d\n", len(s2), cap(s2))
	fmt.Printf("len(s3): %d, cap(s3): %d\n", len(s3), cap(s3))
	fmt.Printf("len(s4): %d, cap(s4): %d\n", len(s4), cap(s4))
	fmt.Printf("len(s5): %d, cap(s5): %d\n", len(s5), cap(s5))
}
