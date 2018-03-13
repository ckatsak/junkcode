package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	WORKLOAD = 300
	WORKERS  = 20
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

		fmt.Printf("[worker %d]: got workload %d! working with it...\n", id, i)

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
	for t := range time.Tick(2 * time.Second) {
		fmt.Printf("[time-waster]: +%v tick...\n", t.Sub(start))
	}
}
