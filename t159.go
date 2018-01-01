package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for range time.Tick(1000 * time.Millisecond) {
		fmt.Printf("%4d.\t%v\n", i, time.Now().Format(time.RubyDate))
		i++
	}
}
