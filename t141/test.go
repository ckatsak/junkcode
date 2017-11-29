package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "test.py")

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// Start the process
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	//time.Sleep(300 * time.Second)
	fmt.Println("exiting")
}
