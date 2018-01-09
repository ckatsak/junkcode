package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	RENAME_NOREPLACE uintptr = 1 << iota
	RENAME_EXCHANGE
	RENAME_WHITEOUT
)

func renameat2(olddirfd uintptr, oldpath string, newdirfd uintptr, newpath string, flags uintptr) error {
	_p0, err := unix.BytePtrFromString(oldpath)
	if err != nil {
		return err
	}
	_p1, err := unix.BytePtrFromString(newpath)
	if err != nil {
		return err
	}
	_, _, er := unix.Syscall6(
		unix.SYS_RENAMEAT2,
		olddirfd,
		uintptr(unsafe.Pointer(_p0)),
		newdirfd,
		uintptr(unsafe.Pointer(_p1)),
		flags,
		0,
	)
	if er == 0 {
		return nil
	}
	return er
}

func test() string {
	// Get cwd.
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("os.Getwd: %v\n", err)
	}
	// Get target file name.
	newAbsPath := filepath.Join(cwd, "Z")
	if !filepath.IsAbs(newAbsPath) {
		return fmt.Sprintf("New file %q should be absolute path!\n", newAbsPath)
	}

	// Create a temp file.
	tmpFile, err := ioutil.TempFile(cwd, "")
	if err != nil {
		return fmt.Sprintf("ioutil.TempFile(): %v\n", err)
	}
	defer func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			log.Printf("ERROR removing %q: %v\n", tmpFile.Name(), err)
		} else {
			log.Printf("Temp file %q removed successfully.\n", tmpFile.Name())
		}
	}()
	tmpAbsPath, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		return fmt.Sprintf("ERROR getting absolute path for %q: %v\n", tmpFile.Name(), err)
	}
	log.Printf("Temp file %q (abs: %q) ready.\n", tmpFile.Name(), tmpAbsPath)

	// Write data to the temp file.
	if err = ioutil.WriteFile(tmpAbsPath, []byte(fmt.Sprintf("CHRISTOULAS @ %s", tmpAbsPath)), 0666); err != nil {
		return fmt.Sprintf("ERROR writing to temp file %q: %v\n", tmpAbsPath, err)
	} else {
		log.Printf("Writing to temp file %q was successful.\n", tmpAbsPath)
	}
	// fdatasync(2) temp file to the disk.
	if err = unix.Fdatasync(int(tmpFile.Fd())); err != nil {
		return fmt.Sprintf("ERROR fdatasync(2) %q: %v\n", tmpAbsPath, err)
	} else {
		log.Printf("fdatasync(2) to temp file %q was successful.\n", tmpAbsPath)
	}

	// renameat2(2) temp file to something else.
	if err = renameat2(0, tmpAbsPath, 0, newAbsPath, RENAME_NOREPLACE); err != nil {
		return fmt.Sprintf("FAILURE: renameat2( . . . , RENAME_NOREPLACE): %v\n", err)
	} else {
		log.Printf("SUCCESS renameat2( . . . , RENAME_NOREPLACE)'ing!")
	}

	// Get cwd's fd.
	cwdF, err := os.Open(cwd)
	if err != nil {
		return fmt.Sprintf("ERROR open(2)'ing %q: %v\n", cwd, err)
	} else {
		log.Printf("open(2) on %q was successful.\n", cwd)
	}
	// fsync(2) cwd's entries to the disk.
	if err = cwdF.Sync(); err != nil {
		return fmt.Sprintf("ERROR fsync(2)'ing %q: %v\n", cwd, err)
	} else {
		log.Printf("fsync(2) on %q was successful.\n", cwd)
	}
	// close(2) cwd's fd.
	if err = cwdF.Close(); err != nil {
		return fmt.Sprintf("ERROR close(2)'ing %q: %v\n", cwd, err)
	} else {
		log.Printf("close(2) on %q was successful.\n", cwd)
	}

	return "NORMAL RETURN"
}

func main() {
	log.Println(test())
}
