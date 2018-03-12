package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	WORKLOAD = 3
	WORKERS  = 2
)

func main() {
	go waste(time.Now())

	ch := make(chan int, WORKLOAD)
	for i := 0; i < WORKLOAD; i++ {
		ch <- i
	}

	wg := new(sync.WaitGroup)
	wg.Add(WORKERS)
	for i := 0; i < WORKERS; i++ {
		go work(i, ch, wg)
	}
	wg.Wait()

	fmt.Printf("[main]: THE END\n")
}

func work(id int, ch chan int, wg *sync.WaitGroup) {
	fmt.Printf("[worker %d]: just spawned!\n", id)
	reEnqueued := make([]bool, WORKLOAD)
	for {
		fmt.Printf("[worker %d]: entering loop...\n", id)

		if len(ch) == 0 {
			fmt.Printf("[worker %d]: breaking from loop...\n", id)
			wg.Done()
			break
		}
		i := <-ch

		t := (1000 + time.Duration(rand.Int31n(2000))) * time.Millisecond
		fmt.Printf("[worker %d]: got workload %d! working with it for %v...\n", id, i, t)
		time.Sleep(t)

		if !reEnqueued[i] {
			reEnqueued[i] = true
			fmt.Printf("[worker %d]: re-enqueuing %d...\n", id, i)
			ch <- i
			continue
		}
		fmt.Printf("[worker %d]: removing %d from the queue!\n", id, i)
	}
	fmt.Printf("[worker %d]: exiting!\n", id)
}

func waste(start time.Time) {
	for t := range time.Tick(3 * time.Second) {
		fmt.Printf("[time-waster]: +%v tick...\n", t.Sub(start))
	}
}
