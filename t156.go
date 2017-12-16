package main

import (
	"fmt"
	"time"
)

const (
	NumInts       = 37
	NumGoroutines = 10
)

func main() {
	c := make(chan int, NumInts)
	for i := 0; i < NumInts; i++ {
		c <- i
	}
	close(c)

	done := make(chan struct{})
	for i := 0; i < NumGoroutines; i++ {
		go func(id int) {
			for x := range c {
				fmt.Printf("%d: %d\n", id, x)
				<-time.After(250 * time.Millisecond)
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < NumGoroutines; i++ {
		<-done
	}
	fmt.Println("main: exiting...")
}
