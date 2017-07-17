package sha256

/*
#include <stdlib.h>
#include <openssl/sha.h>
#cgo LDFLAGS: -lcrypto
*/
import "C"
import (
	//"fmt"
	"unsafe"
)

func cgoSHA256() {
	md := (*C.char)(unsafe.Pointer(C.SHA256((*C.uchar)(unsafe.Pointer(C.CString(msg))),
		C.size_t(len(msg)), (*C.uchar)(nil))))

	// Print:
	//res := C.GoString(md)
	//fmt.Printf("%x", res) // 604e40d7814621e2412a42b8d04e37d56bdea4628f5d622fb782bb644be8722f

	// Benchmark:
	_ = C.GoString(md)
}
