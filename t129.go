package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	bs := make([]byte, 4)
	binary.PutVarint(bs, 42)
	fmt.Println(bs, string(bs))

	fmt.Println([]byte(strconv.Itoa(42)))
}
