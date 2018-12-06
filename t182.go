package main

func m() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	go func() {
		c1 <- 1
		c2 <- 1
	}()
	select {
	case <-c1:
	case <-c2:
		println("no way")
	default:
		println("yo")
	}
}

func main() {
	for i := 0; i < 1000000; i++ {
		m()
	}
}
