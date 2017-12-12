package main

import (
	"fmt"
	"net"
)

func main() {
	// SRV-Record _xmpp-client._tcp.google.com
	cname, addrs, err := net.LookupSRV("xmpp-client", "tcp", "google.com.")
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	fmt.Printf("SRV:\n\tcname: %s\n", cname)

	targets := []string{}
	for i, addr := range addrs {
		fmt.Printf("\n\taddrs[%d].Target\t= %v\n", i, addr.Target)
		fmt.Printf("\taddrs[%d].Port\t= %v\n", i, addr.Port)
		fmt.Printf("\taddrs[%d].Priority = %v\n", i, addr.Priority)
		fmt.Printf("\taddrs[%d].Weight\t= %v\n", i, addr.Weight)
		targets = append(targets, addr.Target)
	}

	fmt.Println("\nTargets:", targets)
	for _, target := range targets {
		addrs, err := net.LookupHost(target)
		if err != nil {
			fmt.Printf("ERR: %v\n", err)
		} else {
			fmt.Printf("\n\t%v\t--> %v", target, addrs[0])
		}
	}
}
