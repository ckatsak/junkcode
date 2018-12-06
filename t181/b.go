package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/blake2b"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
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

	return string(b)
}

var (
	out32 [32]byte
	out20 [20]byte
)

func main() {
	obj := map[string]string{
		"128 B":   RandStringBytesMaskImprSrc(1 << 7),
		"1 KiB":   RandStringBytesMaskImprSrc(1 << 10),
		"512 KiB": RandStringBytesMaskImprSrc(1 << 19),
		"4 MiB":   RandStringBytesMaskImprSrc(1 << 22),
		"32 MiB":  RandStringBytesMaskImprSrc(1 << 25),
	}
	results := map[string]testing.BenchmarkResult{}

	fmt.Println(strings.Repeat("-", 107))
	for name, data := range obj {
		// blake2b
		key := fmt.Sprintf("%27s - %s", "golang.org/x/crypto/blake2b", name)
		objReader := strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out32 = blake2b.Sum256(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// sha1
		key = fmt.Sprintf("%27s - %s", "crypto/sha1", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out20 = sha1.Sum(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		fmt.Println(strings.Repeat("-", 107))
	}
}
