package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	produced  = 0
	processed = 0
)

func produceEndlessly(out chan int, quit chan bool) {
	defer close(out)
	for {
		select {
		case <-quit:
			fmt.Println("RECV QUIT")
			return
		default:
			out <- rand.Int()
			time.Sleep(time.Duration(rand.Int63n(5e6)))
			produced++
		}
	}
}

func quitRandomly(quit chan bool) {
	d := time.Duration(rand.Int63n(5e9))
	fmt.Println("SLEEP", d)
	time.Sleep(d)
	fmt.Println("SEND QUIT")
	quit <- true
}

func main() {
	vals, quit := make(chan int, 10), make(chan bool)
	go produceEndlessly(vals, quit)
	go quitRandomly(quit)
	for {
		x, ok := <-vals
		if !ok {
			break
		}
		fmt.Println(x)
		processed++
		time.Sleep(time.Duration(rand.Int63n(5e8)))
	}
	fmt.Println("Produced:", produced)
	fmt.Println("Processed:", processed)
}
