package main

import (
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

func main() {
	obj := RandStringBytesMaskImprSrc(1 << 22)
	fmt.Println("Object size:", len(obj), "B")
	r1, r2 := strings.NewReader(obj), strings.NewReader(obj)
	fmt.Printf("golang.org/x/crypto/blake2b\t: %x\ngithub.com/minio/blake2b-simd\t: %x\n",
		blake2b4MiB(r1), minio4MiB(r2))
	fmt.Println(`---------------------------------------------------------------------------------------------------`)

	objReader := strings.NewReader(obj)
	brBlake2b := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = blake2b4MiB(objReader)
		}
	})

	objReader = strings.NewReader(obj)
	brMinio := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = minio4MiB(objReader)
		}
	})
	fmt.Println("golang.org/x/crypto/blake2b\t", brBlake2b, brBlake2b.MemString())
	fmt.Println(`---------------------------------------------------------------------------------------------------`)
	fmt.Println("github.com/minio/blake2b-simd\t", brMinio, brMinio.MemString())
}

func blake2b4MiB(reader *strings.Reader) [blake2b.Size256]byte {
	if data, err := ioutil.ReadAll(reader); err != nil {
		panic(err)
	} else {
		return blake2b.Sum256(data)
	}
}

func minio4MiB(reader *strings.Reader) [32]byte {
	if data, err := ioutil.ReadAll(reader); err != nil {
		panic(err)
	} else {
		return minio.Sum256(data)
	}
}
