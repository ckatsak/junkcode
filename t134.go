package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

const msg = "CHRISTOULAS_OEO"

func main() {
	d := blake2b.Sum256([]byte(msg))
	fmt.Printf("%%s d:\t\t%s\n", d)
	fmt.Printf("string(d[:]):\t%v\n", string(d[:]))
	fmt.Printf("%%v d:\t\t%v\n", d)
	fmt.Printf("%%x d:\t\t%x\n", d)
	fmt.Printf("len(d):\t\t%d\n", len(d))
	fmt.Printf("len(%%s of d):\t%d\n", len(fmt.Sprintf("%s", d)))
	fmt.Printf("len(%%x of d):\t%d\n", len(fmt.Sprintf("%x", d)))
}
