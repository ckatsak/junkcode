package main

import (
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	RENAME_NOREPLACE = 1 << iota
	RENAME_EXCHANGE
	RENAME_WHITEOUT
)

func renameat2(olddirfd uintptr, oldpath string, newdirfd uintptr, newpath string, flags int) (err error) {
	_p0, err := unix.BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	_p1, err := unix.BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, err = unix.Syscall6(
		unix.SYS_RENAMEAT2,
		uintptr(olddirfd),
		uintptr(unsafe.Pointer(_p0)),
		uintptr(newdirfd),
		uintptr(unsafe.Pointer(_p1)),
		uintptr(flags),
		0,
	)
	return
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd: %v\n", err)
	}

	oldPath := filepath.Join(cwd, "A")
	newPath := filepath.Join(cwd, "B")
	log.Printf("Attempting to renameat2(2) %q to %q...\n", oldPath, newPath)

	cwdF, err := os.Open(cwd)
	if err != nil {
		log.Fatalf("os.Open(%q): %v\n", cwd, err)
	}

	if err = renameat2(cwdF.Fd(), oldPath, cwdF.Fd(), newPath, RENAME_NOREPLACE); err != nil {
		log.Fatalf("FAILURE: renameat2(): %v\n", err)
	}
	log.Println("SUCCESS")
}
