package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:\n\t$ %s <bytes_2exp>\n", os.Args[0])
	}
	numBytes, err := strconv.ParseUint(os.Args[1], 10, 5)
	if err != nil {
		log.Fatalf("error converting %q to uint: %v\n", os.Args[1], err)
	}

	urfd, err := unix.Open("/dev/urandom", unix.O_RDONLY, 0666)
	if err != nil {
		log.Panicf("unix.Open(): %v\n", err)
	}

	buf := make([]byte, 1<<numBytes)
	start := time.Now()
	if _, err = unix.Read(urfd, buf); err != nil {
		log.Panicf("unix.Read(): %v\n", err)
	}
	end := time.Since(start)

	log.Printf("[ %d B ] in %v\n", len(buf), end)
	//log.Printf("[ %d B ] in %v\n%x\n", len(buf), end, buf)
}
