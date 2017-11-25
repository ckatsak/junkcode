package main

import (
	"bytes"
	"fmt"
	"sort"

	"golang.org/x/crypto/blake2b"
)

type Msg string

const (
	msgM Msg    = `CHRISTOULAS_OEO`
	msgS string = `CHRISTOULAS_OEO`
	msgZ string = "CHRISTOULAS_OEO"
)

func main() {
	fmt.Println(`----------------------------------------------------------------`)
	fmt.Printf("%x\n", blake2b.Sum256([]byte(msgM)))
	fmt.Printf("%x\n", blake2b.Sum256([]byte(msgS)))
	fmt.Printf("%x\n", blake2b.Sum256([]byte(msgZ)))
	fmt.Println(`----------------------------------------------------------------`)

	fmt.Println(`----------------------------------------------------------------`)
	d1 := blake2b.Sum256([]byte(msgM))
	d2 := blake2b.Sum256([]byte(msgS))
	if d1 == d2 {
		fmt.Println("d1 == d2")
	} else {
		fmt.Println("d1 ?= d2")
	}
	fmt.Println(`----------------------------------------------------------------`)

	fmt.Println(`----------------------------------------------------------------`)
	s := []node{
		node{
			blake2b.Sum256([]byte("a")),
		},
		node{
			blake2b.Sum256([]byte("b")),
		},
		node{
			blake2b.Sum256([]byte("c")),
		},
	}
	fmt.Println("Before sorting:", s)
	sort.Slice(s, func(i, j int) bool {
		if bytes.Compare(s[i].dgst[:], s[j].dgst[:]) < 0 {
			return true
		}
		return false
	})
	fmt.Println("After sorting:", s)
	fmt.Println(`----------------------------------------------------------------`)
}

type node struct {
	dgst [blake2b.Size256]byte
}

func (n node) String() string {
	return fmt.Sprintf("%x", n.dgst)
}
