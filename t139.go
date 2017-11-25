package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/blake2b"
)

const msg = `agent-36`

func main() {
	// db is a [32]byte
	db := blake2b.Sum256([]byte(msg))
	fmt.Printf("%x\n", db)
	fmt.Println(db)

	// zb is a *big.Int
	zb := new(big.Int).SetBytes(db[:])
	fmt.Println(zb)

	// hd is a string
	hd := hex.EncodeToString(db[:])
	// hb is a *big.Int
	hb, success := new(big.Int).SetString(hd, 16)
	if !success {
		panic("failure")
	}
	fmt.Println(hb)

	if zb.String() != hb.String() {
		panic("weird")
	} else {
		fmt.Println("so equal")
	}

	zb42 := zb.Add(zb, big.NewInt(int64(42)))
	fmt.Println(zb42, "(increased by 42)")
	fmt.Println(zb42.Bytes(), "(increased by 42)")
	fmt.Printf("%x (increased by 42)\n", zb)

	zb42s := fmt.Sprintf("%064x", zb42)
	fmt.Println(zb42s, len(zb42s))
}
