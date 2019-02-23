package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\n\t./t184 <path>\n\n")
		os.Exit(1)
	}

	fmt.Printf("filepath.Base(%q): %q\n", os.Args[1], filepath.Base(os.Args[1]))
	fmt.Printf("filepath.Dir(%q): %q\n", os.Args[1], filepath.Dir(os.Args[1]))

	abs, err := filepath.Abs(filepath.Dir(os.Args[1]))
	if err != nil {
		panic(err)
	}
	fmt.Printf("filepath.Abs(filepath.Dir(%q)): %q\n", os.Args[1], abs)
}
