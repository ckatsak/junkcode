package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandBytesMaskImprSrc(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:\n$ %s <bytes_2exp>\n", os.Args[0])
	}
	numBytes, err := strconv.ParseUint(os.Args[1], 10, 5)
	if err != nil {
		log.Fatalf("error converting %q to uint: %v\n", os.Args[1], err)
	}

	start := time.Now()
	buf := RandBytesMaskImprSrc(1 << numBytes)
	end := time.Since(start)

	log.Printf("[ %d B ] in %v\n", len(buf), end)
}
