package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golang/glog"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	r := strings.NewReader("")
	defer fmt.Println("done")
	if _, err := r.Read(make([]byte, 1)); err != nil {
		glog.Fatal(err)
	}
}
