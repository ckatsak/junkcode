package main

import "fmt"

type myUint32 uint32

func main() {
	x := myUint32(42)
	incuint32((*uint32)(&x))
	fmt.Printf("%d\n", x)
}

func incuint32(u32 *uint32) {
	*u32 += 10
}
