// Run:
//	$ go run t144.go --logtostderr -v 2
//	$ go run t144.go --logtostderr -v 42
// etc...
package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Infoln("yo")
	if glog.V(1) {
		glog.Infoln(1)
		glog.Warningln(1)
		glog.Errorln(1)
	}
	if glog.V(2) {
		glog.Infoln(2)
		glog.Warningln(2)
		glog.Errorln(2)
	}
	if glog.V(3) {
		glog.Infoln(3)
		glog.Warningln(3)
		glog.Errorln(3)
	}
	time.Sleep(2 * time.Second)
	if glog.V(8) {
		glog.Infoln(8)
		glog.Warningln(8)
		glog.Errorln(8)
	}
	if glog.V(10) {
		glog.Infoln(10)
		glog.Warningln(10)
		glog.Errorln(10)
	}
	if glog.V(42) {
		glog.Infoln(42)
		glog.Warningln(42)
		glog.Errorln(42)
	}
	glog.Infof("%#v\n", glog.Stats)
}
