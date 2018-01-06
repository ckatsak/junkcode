package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	produced  = 0
	processed = 0
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	defer log.Println("[consumer]: Am I out after the producer?!")
	vals, quit := make(chan int, 1), make(chan bool)

	go func() {
		defer log.Println("[producer]: Am I out before the consumer?!")
		defer close(vals)
		i := 0
		for {
			i++
			select {
			case <-quit:
				log.Println("[producer]: RECV QUIT")
				return
			default:
				vals <- i
				produced++
			}
		}
	}()

	//go quitRandomly(quit)
	for {
		x, ok := <-vals
		if !ok {
			break
		}

		// XXX v
		if x == 6 {
			quit <- true
			break
		}
		// XXX ^

		log.Println("[consumer]:", x)
		processed++
		time.Sleep(time.Duration(rand.Int63n(9e8)))
	}
	log.Println("[consumer]: Produced:", produced)
	log.Println("[consumer]: Processed:", processed)
}
