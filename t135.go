package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	limit := 100 + rand.Intn(900)
	for i := 100; i < limit; i++ {
		fmt.Printf("\rProgress status: %.2f %%", float64(i)/float64(limit)*100.0)
		sleepTime := time.Duration(100+rand.Intn(900)) * time.Millisecond
		time.Sleep(sleepTime)
	}
	fmt.Println()
}
