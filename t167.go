package main

import (
	"crypto/sha256"
	"fmt"
)

const PHRASE = `CHRISTOULAS`

func main() {
	hashFunc := func(in []byte) []byte {
		out := sha256.Sum256(in)
		return out[:]
	}

	d := hashFunc([]byte(PHRASE))
	fmt.Printf("d: %x\nd: %T\nlen: %d\ncap: %d\n", d, d, len(d), cap(d))
}
