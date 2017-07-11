package main

import (
	"crypto/sha256"
	"fmt"
)

const (
	testString = "Hello, world!\n"
)

func main() {
	h := sha256.New()
	h.Write([]byte(testString))
	//d := h.Sum(nil)
	d := h.Sum([]byte("WOOHOO"))
	fmt.Printf("len(d) == %d\n\"%%v\": %v\n\"%%s\": %s\n\"%%x\": %x\n", len(d), d, d, d)
}
