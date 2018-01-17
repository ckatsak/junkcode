package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	ls(".")
	ls("skatare")
}

func ls(dir string) {
	fmt.Printf("%s\n$ ls -l %s\n", strings.Repeat("-", 79), dir)
	lsCmd := exec.Command("ls", "-l", dir)
	output, err := lsCmd.CombinedOutput()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			log.Fatalf("returned: %#v\n", exitErr)
		}
	}
	//outStr := string(output)
	//fmt.Printf("%s\n", string(output))
	outBytLines := bytes.Split(output, []byte{'\n'})
	for i, line := range outBytLines {
		fmt.Printf("line %d: %q\n", i, line)

		// get last using bytes
		lineBytesList := bytes.Split(line, []byte{' '})
		fmt.Printf("\tlineBytesList = %q\n", lineBytesList)
		lineBytesLast := lineBytesList[len(lineBytesList)-1]
		fmt.Printf("\tlineBytesLast = %q\n", lineBytesLast)

		// get last using strings
		lineStringList := strings.Split(string(line), " ")
		fmt.Printf("\tlineStringList = %q\n", lineStringList)
		lineStringLast := lineStringList[len(lineStringList)-1]
		fmt.Printf("\tlineStringLast = %q\n", lineStringLast)

		fmt.Println()
	}

	firstOutputLine := bytes.SplitN(output, []byte{'\n'}, 2)[0]
	fmt.Printf("first output line: %q\n", firstOutputLine)
}
