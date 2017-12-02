package main

import (
	"log"
	"net"

	"./customer"
	"google.golang.org/grpc"
)

const port = ":54242"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	customer.RegisterCustomerServer(s, &customer.Server{})
	s.Serve(lis)
}
