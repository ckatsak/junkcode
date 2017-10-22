package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--------------------------------------------------------")

	line := "   <-front: 3 spaces, back: 5 spaces->     "
	fmt.Printf("line = %q\n", line)
	fmt.Println("--------------------------------------------------------")

	line = strings.TrimSpace(line)
	fmt.Printf("line = %+q\n", line)
	fmt.Println("--------------------------------------------------------")
}
