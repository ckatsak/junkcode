package sha256

import "testing"

func BenchmarkGoSHA256(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goSHA256()
	}
}
