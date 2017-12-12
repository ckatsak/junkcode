package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"os/exec"

	"github.com/golang/glog"
)

// Somewhat based on https://stackoverflow.com/a/25191479/2304215.
func execute(cmd ...string) {
	ls := exec.Command(cmd[0], cmd[1:]...)
	stdout, err := ls.StdoutPipe()
	if err != nil {
		glog.Fatal(err)
	}
	stderr, err := ls.StderrPipe()
	if err != nil {
		glog.Fatal(err)
	}
	output := io.MultiReader(stdout, stderr)

	if err := ls.Start(); err != nil {
		glog.Fatal(err)
	}

	in := bufio.NewScanner(output)
	for in.Scan() {
		glog.Infof(in.Text())
	}
	if err := in.Err(); err != nil {
		glog.Error(err)
	}
}

func main() {
	os.Args = append(os.Args, "--logtostderr", "-v=42") // nvm ugliness
	flag.Parse()

	execute("ls")
	execute("ls", "nonexistingdir/")
}
