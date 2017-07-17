package sha256

import "testing"

func BenchmarkCgoSHA256(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cgoSHA256()
	}
}
