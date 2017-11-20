package main

import (
	"log"
	"net"
)

func IsPrivate(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}
	_, p24, _ := net.ParseCIDR("10.0.0.0/8")
	_, p20, _ := net.ParseCIDR("172.16.0.0/12")
	_, p16, _ := net.ParseCIDR("192.168.0.0/16")
	return p24.Contains(ip) || p20.Contains(ip) || p16.Contains(ip)
}

func main() {
	localAddrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}
	for _, localAddr := range localAddrs {
		log.Printf("%s\n", localAddr.String())
		//ip, ipNet, err := net.ParseCIDR(localAddr.String())
		ip, _, err := net.ParseCIDR(localAddr.String())
		if err != nil {
			log.Printf("\tnet.ParseCIDR(): %v\n\n", err)
		}

		if ip.DefaultMask() == nil {
			log.Printf("\t--> is IPv6 address\n")
		} else {
			log.Printf("\t--> is IPv4 address\n")
		}
		if ip.IsLoopback() {
			log.Printf("\t--> is loopback\n")
		}
		if ip.IsMulticast() {
			log.Printf("\t--> is multicast\n")
		}
		if ip.IsUnspecified() {
			log.Printf("\t--> is unspecified\n")
		}
		if ip.IsGlobalUnicast() {
			log.Printf("\t--> is global unicast\n")
		}
		if IsPrivate(ip.String()) {
			log.Printf("\t--> is private <--\n")
		}
		if ip.IsLinkLocalUnicast() {
			log.Printf("\t--> is link-local unicast\n")
		}
		if ip.IsLinkLocalMulticast() {
			log.Printf("\t--> is link-local multicast\n")
		}
		if ip.IsInterfaceLocalMulticast() {
			log.Printf("\t--> is interface-local multicast\n")
		}
	}
}
