package main

import "fmt"

func main() {
	msg := `CHRISTOULAS`
	fmt.Println(msg[:1])
	fmt.Println(msg[:3])
	fmt.Println(msg[:5])

	fmt.Println(msg[1])
	fmt.Println(string(msg[1]))
	fmt.Println(msg[3])
	fmt.Println(string(msg[3]))
	fmt.Println(msg[5])
	fmt.Println(string(msg[5]))

	printUint32(0666)
}

func printUint32(x uint32) {
	fmt.Printf("%#b == %#o == %#d == %#x\n", x, x, x, x)
}
