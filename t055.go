//
// Goroutines basics
//
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	NR_ROUTINES = 3
	SLEEP_LIMIT = 5
	WATCH_FRAME = 60
)

//
func main() {
	for i := 0; i < NR_ROUTINES; i++ {
		go test(i)
	}
	time.Sleep(WATCH_FRAME * time.Second)

	return
}

//
func test(i int) {
	for {
		t := time.Duration(rand.Int31n(SLEEP_LIMIT))
		time.Sleep(t * time.Second)
		fmt.Fprintf(os.Stderr, "Hello from goroutine [%d], after %d sec of sleep\n", i, int(t))
	}
}
