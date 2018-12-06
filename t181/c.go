package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/cespare/xxhash"
	"github.com/dchest/siphash"
	"github.com/minio/highwayhash"
	"github.com/spaolacci/murmur3"
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
	out32        [32]byte
	out20        [20]byte
	out16        [16]byte
	out16s       []byte = make([]byte, 16)
	out8a, out8b uint64
)

func main() {
	obj := map[string]string{
		"128 B":   RandStringBytesMaskImprSrc(1 << 7),
		"1 KiB":   RandStringBytesMaskImprSrc(1 << 10),
		"128 KiB": RandStringBytesMaskImprSrc(1 << 17),
		"512 KiB": RandStringBytesMaskImprSrc(1 << 19),
		"1 MiB":   RandStringBytesMaskImprSrc(1 << 20),
		"2 MiB":   RandStringBytesMaskImprSrc(1 << 21),
		"4 MiB":   RandStringBytesMaskImprSrc(1 << 22),
		"32 MiB":  RandStringBytesMaskImprSrc(1 << 25),
	}
	results := map[string]testing.BenchmarkResult{}

	fmt.Println(strings.Repeat("-", 108))
	for name, data := range obj {
		// blake2b
		key := fmt.Sprintf("%28s - %s", "golang.org/x/crypto/blake2b", name)
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
		key = fmt.Sprintf("%28s - %s", "crypto/sha1", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out20 = sha1.Sum(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// md5
		key = fmt.Sprintf("%28s - %s", "crypto/md5", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out16 = md5.Sum(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// keyed hash functions' key
		hashKey128 := make([]byte, 16)
		hashKey256 := make([]byte, 32)

		// SipHash
		key = fmt.Sprintf("%28s - %s", "github.com/dchest/siphash", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out16s = siphash.New128(hashKey128).Sum(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// HighwayHash
		key = fmt.Sprintf("%28s - %s", "github.com/minio/highwayhash", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out16 = highwayhash.Sum128(data, hashKey256)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// MurmurHash3
		key = fmt.Sprintf("%28s - %s", "github.com/spaolacci/murmur3", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out8a, out8b = murmur3.Sum128(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		// xxHash
		key = fmt.Sprintf("%28s - %s", "github.com/cespare/xxhash", name)
		objReader = strings.NewReader(data)
		results[key] = testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				data, _ := ioutil.ReadAll(objReader)
				out8a = xxhash.Sum64(data)
			}
		})
		fmt.Println(key, "\t", results[key], results[key].MemString())

		fmt.Println(strings.Repeat("-", 108))
	}
}
