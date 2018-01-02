package main

import (
	"log"
	"time"
)

const (
	waitTime = 1 * time.Second
	NUM      = 3
)

func yolo(stopCh <-chan struct{}) <-chan int {
	retCh := make(chan int)
	go func() {
		defer log.Println(`[worker]: Leaving now!`)
		defer close(retCh)
		i := 0
		for {
			i++
			select {
			case <-stopCh:
				log.Println(`[worker]: Leaving this world soon...`)
				retCh <- -42
				return
			case retCh <- i:
				//default:
				//	retCh <- i
			}
		}
	}()
	return retCh
}

func main() {
	stopCh := make(chan struct{})
	ints := yolo(stopCh)
	for j := 0; j < NUM; j++ {
		log.Println("[main]: Got:", <-ints)
		log.Println("[main]: Calculating...")
		time.Sleep(waitTime)
	}
	close(stopCh)
	log.Println("[main]: More calculations now...")
	time.Sleep(time.Nanosecond)
	log.Println(`[main]: Should receive -42 now...`)
	log.Printf("[main]: Is %d == -42 ?\n", <-ints)
	log.Println(`[main]: Further attempts to read from the chan should return 0 (because it should be closed)...`)
	log.Printf("[main]: Is %d == 0 ?\n", <-ints)
}
