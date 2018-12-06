package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
	"time"

	minio "github.com/minio/blake2b-simd"
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

func genericHash(reader *strings.Reader, hashFunc func([]byte) [32]byte) [32]byte {
	if data, err := ioutil.ReadAll(reader); err != nil {
		panic(err)
	} else {
		return hashFunc(data)
	}
}

func sha1Sum32Aux(in []byte) (out [32]byte) {
	tmp := sha1.Sum(in)
	copy(out[:], tmp[:])
	return
}

func main() {
	obj := map[string]string{
		"128 B":   RandStringBytesMaskImprSrc(1 << 7),
		"1 KiB":   RandStringBytesMaskImprSrc(1 << 10),
		"512 KiB": RandStringBytesMaskImprSrc(1 << 19),
		"4 MiB":   RandStringBytesMaskImprSrc(1 << 22),
		"32 MiB":  RandStringBytesMaskImprSrc(1 << 25),
	}
	hash := map[string]func([]byte) [32]byte{
		"golang.org/x/crypto/blake2b":   blake2b.Sum256,
		"github.com/minio/blake2b-simd": minio.Sum256,
		"crypto/sha256":                 sha256.Sum256,
		"crypto/sha512":                 sha512.Sum512_256,
		"crypto/sha1":                   sha1Sum32Aux,
	}
	results := map[string]testing.BenchmarkResult{}

	fmt.Println(strings.Repeat("-", 115))
	for name, data := range obj {
		for pkg, hsh := range hash {
			key := fmt.Sprintf("%29s - %s", pkg, name)
			objReader := strings.NewReader(data)
			results[key] = testing.Benchmark(func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					genericHash(objReader, hsh)
				}
			})
			fmt.Println(key, "\t", results[key], results[key].MemString())
		}
		fmt.Println(strings.Repeat("-", 115))
	}
}
