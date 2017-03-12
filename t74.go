package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func init() {
	os.Setenv("ENV_VAR", "web-42")
}

func main() {
	env := os.Getenv("ENV_VAR")

	var id int
	if i := bytes.LastIndexByte([]byte(env), "-"[0]); i != -1 {
		var err error
		if id, err = strconv.Atoi(env[i+1:]); err != nil {
			fmt.Println("Atoi:", err)
		}
	}

	fmt.Println("env:", env)
	fmt.Println("id:", id)

	return
}
