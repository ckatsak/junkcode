package main

import (
	"fmt"
	"strconv"

	"github.com/minio/highwayhash"
)

func main() {
	for i := 0; i < 35; i++ {
		fmt.Printf("%x\n", highwayhash.Sum128([]byte(strconv.Itoa(i)), make([]byte, 32)))
	}
}
