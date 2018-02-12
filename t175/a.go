package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const msg = "CHRISTOULAS"

// Bitset to string serialization: bit to agentID mapping:
//
//   0  1  0  0  0  0  1  1    0  1  0  0  1  0  0  0    0  1  0  1  0  0  1  0    0  1  0  0  1  0  0  1
//   ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    ^  ^  ^  ^  ^  ^  ^  ^    . . .
//   7  6  5  4  3  2  1  0   15 14 13 12 11 10  9  8   23 22 21 20 19 18 17 16   31 30 29 28 27 26 25 24

func bitIsSet(vector string, id int) bool {
	//fmt.Printf("\n%v\n", vector[id>>3]&(1<<(uint(id)&7)))
	//
	//          if != 0, will be a power of 2
	//      vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	return (vector[id>>3] & (1 << (uint(id) & 7))) != 0
}

func setBit(vector *string, id int) {
	//(*vector)[id>>3] |=
}

func clearBit(vector *string, id int) {
	b := (*vector)[id>>3]
	(*vector)[id>>3] &^= (1 << (uint(id) & 7))
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

	coarseIndex, fineIndex := ID>>3, ID&7
	fmt.Printf("\nID: %d --> coarse: %d, fine: %d\n", ID, coarseIndex, fineIndex)
	fmt.Printf("\nstring[%d]: Type: %T; Content: %q (%d == %08b)\n",
		coarseIndex, msg[coarseIndex], msg[coarseIndex], msg[coarseIndex], msg[coarseIndex])

	fmt.Printf("bitIsSet(%q, %d): %t\n", msg, ID, bitIsSet(msg, ID))
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
	fmt.Printf("string: %s\n", msg)
}

func printBytes(vector string) {
	fmt.Printf("[]byte: %v\n", []byte(msg))
}
