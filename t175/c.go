package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

const msg = "HTANENAMIKROKARAVIPOYHTANATAKSIDEFTOOEOEOEOE"

var emsg = base64.StdEncoding.EncodeToString([]byte(msg))

// Bitset to []byte serialization: bit to agentID mapping:
//
//   0  1  0  0  0  0  1  1    0  1  0  0  1  0  0  0    0  1  0  1  0  0  1  0    0  1  0  0  1  0  0  1
//   ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    . . .
//   7  6  5  4  3  2  1  0   15 14 13 12 11 10  9  8   23 22 21 20 19 18 17 16   31 30 29 28 27 26 25 24

func bitIsSet(encodedVector string, id int) bool {
	bitvector, err := base64.StdEncoding.DecodeString(encodedVector)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("\n%v\n", bitvector[id>>3]&(1<<(uint(id)&7)))
	//
	//            if != 0, will be a power of 2
	//      vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	return (bitvector[id>>3] & (1 << (uint(id) & 7))) != 0
}

func setBit(encodedVector *string, id int) {
	bitvector, err := base64.StdEncoding.DecodeString(*encodedVector)
	if err != nil {
		panic(err)
	}
	bitvector[id>>3] |= 1 << (uint(id) & 7)
	*encodedVector = base64.StdEncoding.EncodeToString(bitvector)
}

func clearBit(encodedVector *string, id int) {
	bitvector, err := base64.StdEncoding.DecodeString(*encodedVector)
	if err != nil {
		panic(err)
	}
	bitvector[id>>3] &^= 1 << (uint(id) & 7)
	*encodedVector = base64.StdEncoding.EncodeToString(bitvector)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n\t$ go run t175.go <agentID>\n")
		os.Exit(1)
	}
	ID, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing <agentID>: %v\n", err)
		os.Exit(1)
	}
	printString(msg)
	printBytes(msg)
	printBitSet(msg)
	//fmt.Printf("\n%s\n", string(bytes.Repeat([]byte{0xff}, 11)))
	//fmt.Printf("%s\n", string(bytes.Repeat([]byte{0x72}, 11)))

	printString(emsg)
	printBytes(emsg)

	coarseIndex, fineIndex := ID>>3, ID&7
	fmt.Printf("\nID: %d --> coarse: %d, fine: %d\n", ID, coarseIndex, fineIndex)
	fmt.Printf("\nstring[%d]: Type: %T; Content: %q (%d == %08b)\n",
		coarseIndex, msg[coarseIndex], msg[coarseIndex], msg[coarseIndex], msg[coarseIndex])

	fmt.Printf("bitIsSet(%q, %d): %t\n", msg, ID, bitIsSet(emsg, ID))

	//
	// FLIP ALL, TWICE:
	//
	fmt.Printf("\nBefore:\n")
	printString(emsg)
	if msg, err := base64.StdEncoding.DecodeString(emsg); err != nil {
		panic(err)
	} else {
		printString(string(msg))
		printBitSet(string(msg))
	}
	for i := 0; i < 2; i++ {
		for id := 0; id < 8*len(msg); id++ {
			if bitIsSet(emsg, id) {
				clearBit(&emsg, id)
			} else {
				setBit(&emsg, id)
			}
		}

		fmt.Printf("\nAfter round %d:\n", i+1)
		printString(emsg)
		if msg, err := base64.StdEncoding.DecodeString(emsg); err != nil {
			panic(err)
		} else {
			printString(string(msg))
			printBitSet(string(msg))
		}
	}
}

func printBitSet(vector string) {
	var buf bytes.Buffer
	//for _, b := range vector {
	//	buf.WriteString(fmt.Sprintf("% 08b", byte(b)))
	//}
	for i := 0; i < len(vector); i++ {
		buf.WriteString(fmt.Sprintf("%08b ", vector[i]))
	}
	fmt.Printf("%q --> %s\n", vector, buf.String())
}

func printString(vector string) {
	fmt.Printf("string: %s\n", vector)
}

func printBytes(vector string) {
	fmt.Printf("[]byte: %v\n", []byte(vector))
}
