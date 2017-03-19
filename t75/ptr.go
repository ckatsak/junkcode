/*
 * Pointer assignment demo
 */
package main

import "fmt"

type str_t struct {
	x, y int
}

type Request struct {
	id   int
	init *int

	buf string

	s1 str_t
	s2 *str_t
}

func (req *Request) PrettyPrintReq(name string) {
	fmt.Printf("\nRequest %s:\n", name)
	fmt.Printf("\taddress: %p\n", req)

	fmt.Printf("\tid: %d\n", req.id)
	fmt.Printf("\tinit: %p\n", req.init)
	fmt.Printf("\t  *init: %d\n", *req.init)

	fmt.Printf("\tbuf:\n")
	fmt.Printf("\t  &buf: %p\n", &req.buf)
	fmt.Printf("\t  buf: \"%s\"\n", req.buf)

	fmt.Printf("\ts1 (str_t):\n")
	fmt.Printf("\t  req.s1.x: %d\n", req.s1.x)
	fmt.Printf("\t  req.s1.y: %d\n", req.s1.y)

	fmt.Printf("\ts2 (*str_t):\n")
	fmt.Printf("\t  s2: %p\n", req.s2)
	if req.s2 != nil {
		fmt.Printf("\t  req.s2.x: %d\n", req.s2.x)
		fmt.Printf("\t  req.s2.y: %d\n", req.s2.y)
	}

	fmt.Println()
}

func main() {
	init_tmp := new(int)
	*init_tmp = 505
	a := &Request{
		id:   42,
		init: init_tmp,
		buf:  "Yo yo yo, world!",
		s1:   str_t{17, 17},
	}
	a.PrettyPrintReq("a")

	b := a
	b.PrettyPrintReq("b (b=a)")

	c := new(Request)
	*c = *a
	c.PrettyPrintReq("c (*c=*a)")

	fmt.Printf("------------------------------------------------------\n")
	fmt.Printf("Modifying a.buf content...\n")
	a.buf = "Modified content!"
	a.PrettyPrintReq("a")
	b.PrettyPrintReq("b (b=a)")
	c.PrettyPrintReq("c (*c=*a)")

	fmt.Printf("------------------------------------------------------\n")
	fmt.Printf("Allocating a.s2 content...\n")
	a.s2 = &str_t{1, 2}
	a.PrettyPrintReq("a")
	b.PrettyPrintReq("b (b=a)")
	c.PrettyPrintReq("c (*c=*a)")

	fmt.Printf("------------------------------------------------------\n")
	fmt.Printf("Modifying *a.init...\n")
	*a.init = 333
	a.PrettyPrintReq("a")
	b.PrettyPrintReq("b (b=a)")
	c.PrettyPrintReq("c (*c=*a)")

	return
}
