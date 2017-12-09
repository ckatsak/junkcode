package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	bins, err := filepath.Glob("/*/*/*/kubectl")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bins)
}
