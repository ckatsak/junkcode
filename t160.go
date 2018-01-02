package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const INDEX_HEIGHT = 4

func main() {
	d1, d2 := sha256.Sum256([]byte("first")), sha256.Sum256([]byte("last"))
	d1[0], d1[1] = 0x4b, 0x20
	d2[0], d2[1] = 0x6a, 0xb0
	f, l := d1[:], d2[:]
	fs, ls := hex.EncodeToString(f), hex.EncodeToString(l)

	fmt.Printf("f: %x\nl: %x\n\n", f, l)
	fmt.Printf("f: %#v\t(#%d)\nl: %#v\t(#%d)\n\n", f, len(f), l, len(l))
	fmt.Printf("f: %s\t(#%d)\nl: %s\t(#%d)\n\n", fs, len(fs), ls, len(ls))

	strSlice := fs[:INDEX_HEIGHT]
	fmt.Printf("fs[:%d] (%T) : %v\n", INDEX_HEIGHT, strSlice, strSlice)
	bytSlice := []byte(fs[:INDEX_HEIGHT])
	fmt.Printf("[]byte(fs[:%d]) (%T) : %#v\n", INDEX_HEIGHT, bytSlice, bytSlice)
	strBytSlice := string([]byte(fs[:INDEX_HEIGHT]))
	fmt.Printf("string([]byte(fs[:%d])) (%T) : %s\n", INDEX_HEIGHT, strBytSlice, strBytSlice)
	decSlice, _ := hex.DecodeString(fs[:INDEX_HEIGHT])
	fmt.Printf("decode(fs[:%d]) (%T) : %#v\n", INDEX_HEIGHT, decSlice, decSlice)

	for i := 1; i <= INDEX_HEIGHT; i++ {
		fmt.Println(strings.Repeat(`-`, 120))
		fmt.Printf("f[:%d]: %x\nl[:%d]: %x\n", i, f[:i], i, l[:i])
		fmt.Printf("fs[:%d]: %s\nls[:%d]: %s\n", i, fs[:i], i, ls[:i])
	}
}
