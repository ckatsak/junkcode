package main

import (
	"fmt"
	"strings"
)

var abc = `abcdefghijklmnopqrstuvwxyz`

func main() {
	fmt.Println(strings.Repeat(`-`, 120))

	// Test strings.Trimspace()
	line := `   <-front: 3 spaces, back: 5 spaces->     `
	fmt.Printf("Initially: line = %q\n", line)
	fmt.Println(strings.Repeat(`-`, 120))

	line = strings.TrimSpace(line)
	fmt.Printf("TrimSpace: line = %+q\n", line)
	fmt.Println(strings.Repeat(`-`, 120))

	// Test strings.Trimleft()
	line = `   <-front: 3 spaces, back: 5 spaces->     `
	line = strings.TrimLeft(line, "\t \n")
	fmt.Printf("TrimLeft: line = %+q\n", line)
	fmt.Println(strings.Repeat(`-`, 120))

	// Test raw strings & strings indexing
	fmt.Printf("abc = (%+q, %T)\n", abc, abc)
	for i := 0; i < len(abc); i += 2 {
		fmt.Printf("abc[%d] = (%+q, %T)\n", i, abc[i], abc[i])
	}
	fmt.Println(strings.Repeat(`-`, 120))

	// Test byte sequences
	abc_bs := []byte(abc)
	fmt.Printf("abc_bs = (%+q, %T)\n", abc_bs, abc_bs)
	fmt.Println(strings.Repeat(`-`, 120))

	// Test range loops over strings
	for i, c := range abc {
		if i%2 == 0 {
			fmt.Printf("abc[%d] = (%+q, %T)\n", i, c, c)
		}
	}
	fmt.Println(strings.Repeat(`-`, 120))
}
