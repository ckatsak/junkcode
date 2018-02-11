package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Printf("Host to lookup:\n>>> ")
	var host string
	if _, err := fmt.Scanf("%s", &host); err != nil {
		fmt.Fprintln(os.Stderr, "error: fmt.Scanf: %v", err)
		os.Exit(1)
	}

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: net.LookupHost: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d results. Showing IPv4 only:\n", len(addrs))
	for i, addr := range addrs {
		if isIPv4 := net.ParseIP(addr).To4(); isIPv4 == nil {
			continue
		}
		fmt.Printf("%d.\t%v\n", i, addr)
	}
}
