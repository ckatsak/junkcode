package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/blake2b"
)

type Node string

func main() {
	for i := -5; i < 5; i++ {
		fmt.Println(i % 10)
	}
	fmt.Println(`----------------------------------------------------------------------`)
	fmt.Println(`----------------------------------------------------------------------`)

	msg := Node("CHRISTOULAS")
	fmt.Printf("%T\n", msg)
	fmt.Printf("%s\n", printString(msg))
	fmt.Println(`----------------------------------------------------------------------`)
	fmt.Println(`----------------------------------------------------------------------`)

	var n Node
	if n == "" {
		fmt.Println("empty string")
	} else {
		fmt.Println(n)
	}
	fmt.Println(`----------------------------------------------------------------------`)
	fmt.Println(`----------------------------------------------------------------------`)

	x := "abcdef"
	fmt.Println("string:", len(x), x)
	fmt.Println("[]byte:", len([]byte(x)), []byte(x))
	if h, err := hex.DecodeString(x); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("hex.DecodeString: %d %v --> %s\n", len(h), h, h)
	}
	fmt.Println(`----------------------------------------------------------------------`)
	fmt.Println(`----------------------------------------------------------------------`)

	dig := blake2b.Sum256([]byte(msg))
	fmt.Printf("%%v(dig %T):\t%v\t%d\n", dig, dig, len(dig))
	hb := hex.EncodeToString(dig[:])
	fmt.Printf("%%s(hb %T)\t\t%s\t%d\n", hb, hb, len(hb))
	fmt.Printf("%%v(hb %T)\t\t%v\t%d\n", hb, hb, len(hb))
	fmt.Printf("%%x(dig %T)\t%x\t%d\n", dig, dig, len(hb))
	fmt.Println(`----------------------------------------------------------------------`)
	if dig2, err := hex.DecodeString(hb); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%%s(dig2 %T)\t%s\t%d\n", dig2, dig2, len(dig2))
		fmt.Printf("%%v(dig2 %T)\t%v\t%d\n", dig2, dig2, len(dig2))
		fmt.Printf("%%x(dig2 %T)\t%x\t%d\n", dig2, dig2, len(dig2))
	}
}

func printString(x Node) string {
	return string(x)
}
