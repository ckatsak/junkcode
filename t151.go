package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	bins, err := filepath.Glob("/*/*/*/kubectl")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bins)

	bin, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bin)
}
