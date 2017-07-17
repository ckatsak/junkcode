package sha256

import (
	"crypto/sha256"
	//"fmt"
)

func goSHA256() {
	h := sha256.New()
	h.Write([]byte(msg))

	// Print:
	//d := h.Sum(nil)
	//fmt.Printf("%x", d)

	// Benchmark:
	_ = h.Sum(nil)
}
