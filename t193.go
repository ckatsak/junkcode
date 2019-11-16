/*
 * Find the index of every occurrence of rune '%' in string.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "100%% σαλάτας %d %%μ%q"
	fmt.Printf("string: %q\n\n", s1)

	s := s1
	i, j, count := 0, 0, 0
	if i = strings.IndexRune(s, '%'); i < 0 {
		fmt.Println("NOPE")
	}

	indices := make([]int, 0)
	for i = 0; i > -1; {
		//fmt.Printf("Current string: %q\ni = %d ; j = %d ; count = %d\n ...\n", s, i, j, count)
		j += i + 1
		i = strings.IndexRune(s, '%')
		//fmt.Printf("i = %d ; j = %d", i, j)
		if i > -1 {
			count++
		} else {
			break
		}

		//fmt.Println("'%' @", i+j-1)
		indices = append(indices, i+j-1)
		s = s1[i+j:]

		//fmt.Println()
	}

	fmt.Println("indices =", indices, "\ncount =", count)
}
