package main

import (
	"fmt"
	"runtime"

	"golang.org/x/crypto/blake2b"
)

type T struct {
	name [32]byte
}

func (t *T) String() string {
	return fmt.Sprintf("%x", t.name)
}

type S struct {
	name []byte
}

func (s *S) String() string {
	return fmt.Sprintf("%x", s.name)
}

func main() {
	// T stuff
	t1 := &T{
		name: blake2b.Sum256([]byte("yo")),
	}
	fmt.Println("1: t1:", t1)

	t2 := &T{
		name: t1.name,
	}
	t1.name = blake2b.Sum256([]byte("hey"))
	fmt.Println("2: t1:", t1)
	fmt.Println("2: t2:", t2)

	// S stuff

	nameS1 := blake2b.Sum256([]byte("yo"))
	s1 := &S{
		name: nameS1[:],
	}
	fmt.Println("3: s1:", s1)
	s2 := &S{
		name: s1.name,
	}
	nameS1 = blake2b.Sum256([]byte("hey"))
	s1.name = nameS1[:]
	fmt.Println("4: s1:", s1)
	fmt.Println("4: s2:", s2)

	s1 = nil
	runtime.GC()
	fmt.Println("5: s2:", s2)
}
