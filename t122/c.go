package main

import (
	"fmt"
	"io"
	"log"
)

func reversed(s string) string {
	n := 0
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	var (
		err    error
		buffer string
	)
	for err != io.EOF {
		fmt.Print("Input: ")
		_, err = fmt.Scanln(&buffer)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Reversed:", reversed(string(buffer)))
	}
}
