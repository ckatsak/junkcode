// getcwd(2), chroot(2), chdir(2) tests
package main

import (
	"fmt"
	"log"
	"syscall"
)

const (
	BUF_SZ = 256
)

func main() {
	cwd1 := make([]byte, BUF_SZ)

	fmt.Print("This is where I am now: ")
	if _, err := syscall.Getcwd(cwd1); err != nil {
		log.Fatal("Error: syscall.Getcwd(2): ", err)
	}
	fmt.Println(string(cwd1))

	fmt.Println("chroot(2)ing back to /home/christos")
	// Commented out doesn't work for some reason. TODO: Why not?
	//cwd_s := string(cwd[:])
	//if err := syscall.Chroot(cwd_s); err != nil {
	if err := syscall.Chroot("/home/christos"); err != nil {
		log.Fatal("Error: syscall.Chroot(2): ", err)
	}

	fmt.Print("This is where I am now: ")
	cwd2 := make([]byte, BUF_SZ)
	if _, err := syscall.Getcwd(cwd2); err != nil {
		log.Fatal("Error: syscall.Getcwd(2): ", err)
	}
	fmt.Println(string(cwd2))

	fmt.Println("chdir(2)ing to parent")
	if err := syscall.Chdir(".."); err != nil {
		log.Fatal("Error: syscall.Chdir(2): ", err)
	}

	fmt.Print("This is where I am now: ")
	cwd3 := make([]byte, BUF_SZ)
	if _, err := syscall.Getcwd(cwd3); err != nil {
		log.Fatal("Error: syscall.Getcwd(2): ", err)
	}
	fmt.Println(string(cwd3))

	fmt.Println("Again, chdir(2)ing to parent")
	if err := syscall.Chdir(".."); err != nil {
		log.Fatal("Error: syscall.Chdir(2): ", err)
	}

	fmt.Print("This is where I am now: ")
	cwd4 := make([]byte, BUF_SZ)
	if _, err := syscall.Getcwd(cwd4); err != nil {
		log.Fatal("Error: syscall.Getcwd(2): ", err)
	}
	fmt.Println(string(cwd4))

	return
}
