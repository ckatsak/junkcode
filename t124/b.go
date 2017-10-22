// Dumb counter of files and directories, given a root directory as a command
// line argument.
//
// Try it on linux kernel source code to see a.go 's BUG solved.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func init() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("")
}

func setrlimit() {
	var rlimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		log.Panicln(err)
	}

	rlimit.Cur = rlimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Rlimit:", rlimit)
}

type Result struct {
	numFiles int
	numDirs  int
}

func main() {
	setrlimit()
	result := Result{}

	if len(os.Args) != 2 {
		log.Fatalf("Usage:\n\t$ %s <dir_path>\n", os.Args[0])
	}
	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	fi, err := os.Stat(root)
	if err != nil {
		log.Panicln(err)
	}
	if fi.IsDir() {
		result = recDir(root)
	} else if fi.Mode().IsRegular() {
		result = fileFunc(root)
	}

	fmt.Println(result)
}

func recDir(root string) (result Result) {
	if filepath.Base(root) == ".git" {
		log.Printf("Skipping %q.", root)
		return
	}
	dir, err := os.Open(root) // open(2)
	if err != nil {
		log.Panicln(err)
	} /*else {
		log.Printf("directory %q\n", root)
	}*/
	defer dir.Close()                // close(2) atexit
	fileinfos, err := dir.Readdir(0) // readdir(2) + stat(2)
	if err != nil {
		log.Panicln(err)
	}

	// Spawn goroutines
	dc := make(chan Result)
	count := 0
	for _, fi := range fileinfos {
		filename := filepath.Join(root, fi.Name())
		if fi.IsDir() {
			count++
			go func(path string) {
				dc <- recDir(path)
			}(filename)
		} else if fi.Mode().IsRegular() {
			count++
			go func(path string) {
				dc <- fileFunc(path)
			}(filename)
		} else {
			log.Printf("%q is alien.\n", filename)
		}
	}

	// Gather the results
	for ; count > 0; count-- {
		r := <-dc
		result.numDirs += r.numDirs
		result.numFiles += r.numFiles
	}
	close(dc)

	result.numDirs++
	return
}

func fileFunc(path string) (result Result) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	} /*else {
		log.Printf("file %q\n", path)
	}*/
	defer file.Close()
	return Result{numFiles: 1}
}
