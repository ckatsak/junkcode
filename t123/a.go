// Uhh... This one just makes no sense.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_GOROUTINES = 1
	MAX_LEVEL      = 3
	MIN_LEVEL      = -MAX_LEVEL
	MAX_SLEEP_SEC  = 2
)

func rec(id, level int) int {
	sleepTime := time.Duration(rand.Int31n(MAX_SLEEP_SEC)) * time.Second
	fmt.Printf("id:%d, level:%d, sleeping for %s\n", id, level, sleepTime)
	time.Sleep(sleepTime)
	if level >= MAX_LEVEL || level <= MIN_LEVEL {
		return level
	}

	// Spawn
	cplus := make(chan int)
	cmult := make(chan int)
	go func() {
		cplus <- rec(id, level+1)
	}()
	go func() {
		cmult <- rec(id, level*level)
	}()

	// Wait
	var plus, mult int = 0, 0
	for i := 0; i < 2; i++ {
		select {
		case p := <-cplus:
			//println(id, p)
			fmt.Printf("id: %d, received p=%d\n", id, p)
			plus = p
		case m := <-cmult:
			//println(id, m)
			fmt.Printf("id: %d, received m=%d\n", id, m)
			mult = m
		}
	}
	close(cplus)
	close(cmult)
	return plus + mult
}

func main() {
	ch := make(chan int)
	count := 0
	for i := 0; i < NUM_GOROUTINES; i++ {
		count++
		go func(id int) {
			ch <- rec(id, 0)
		}(i)
	}

	i := 0
	for v := range ch {
		i++
		println("Main: received", v)
		if i == count {
			break
		}
	}
}
