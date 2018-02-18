package main

import (
	"fmt"
	"time"
)

func main() {
	v := <-yaw()
	fmt.Println(v)
}

func yaw() <-chan error {
	ret := make(chan error)
	go func() {
		time.Sleep(3 * time.Second)
		//ret <- fmt.Errorf("yo")
		close(ret)
	}()
	return ret
}
