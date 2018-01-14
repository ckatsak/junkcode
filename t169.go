package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	d := sha256.Sum256([]byte("skata"))
	fmt.Printf("\t\t%x\n", d)

	dp := new(big.Int).SetBytes(d[:])
	dp.Add(dp, big.NewInt(1))
	fmt.Printf(" + 1 -->\t%x\n", dp)

	//e := new(big.Int).SetBytes(bytes.Repeat([]byte{0xee}, 32))
	//fmt.Printf("\t\t%x\n", e)

	dp.Mod(dp, new(big.Int).SetBytes(bytes.Repeat([]byte{0xdd}, 32)))
	fmt.Printf(" %% ddd...dd -->\t%0[2]*[1]x\n", dp, 64)

	orig := new(big.Int).SetBytes(bytes.Repeat([]byte{0xdd}, 32))
	orig.Add(orig, dp).Sub(orig, big.NewInt(1))
	fmt.Printf(" original -->\t%x\n", orig)
}
