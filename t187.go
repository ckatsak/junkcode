package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, 8)
	fmt.Println(0, buf)
	fmt.Printf("%x\n\n", buf)

	var i uint64 = 513
	binary.BigEndian.PutUint64(buf, i)
	fmt.Println(i, buf)
	fmt.Printf("%x\n\n", buf)

	i = 254
	binary.BigEndian.PutUint64(buf, i)
	fmt.Println(i, buf)
	fmt.Printf("%x\n\n", buf)
}
