package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	ls := exec.Command("ls")
	ls.Stdout = os.Stderr
	ls.Stderr = os.Stderr
	if err := ls.Run(); err != nil {
		log.Fatalln(err)
	}
}
