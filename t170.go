package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	d := sha256.Sum256([]byte("heyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"))
	fmt.Printf("%x\n", d)

	n := new(big.Int).SetBytes(d[:])
	fmt.Printf("n = %x, len(n) = %d\n", n, len(n.Bytes()))

	n2 := new(big.Int).SetBytes(bytes.Repeat([]byte{0xff}, 32))
	fmt.Printf("n2 = %x, len(n2) = %d\n", n2, len(n2.Bytes()))
	n2.Add(n2, big.NewInt(1))
	fmt.Printf("n2 = %x, len(n2) = %d\n", n2, len(n2.Bytes()))

	m := n2.Bytes()
	if len(m) > 32 {
		m = m[1:]
	}
	fmt.Printf("m = %x, len(m) = %d\n", m, len(m))
}
