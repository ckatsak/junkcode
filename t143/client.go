package main

import (
	"context"
	"io"
	"log"

	"./customer"
	"google.golang.org/grpc"
)

const address = "localhost:54242"

func createCustomer(client customer.CustomerClient, customer *customer.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Println("A new Customer has been added with id: %d", resp.Id)
		log.Println("resp.String():", resp.String())
		log.Printf("resp: %v\n", resp)
	}
}

func getCustomers(client customer.CustomerClient, filter *customer.CustomerFilter) {
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v\n", customer)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := customer.NewCustomerClient(conn)

	tux := &customer.CustomerRequest{
		Id:    101,
		Name:  "Tux",
		Email: "tux@linux.com",
		Phone: "42-42-TUX",
		Addresses: []*customer.CustomerRequest_Address{
			&customer.CustomerRequest_Address{
				Street:            "Kifissias 4242",
				City:              "Athens",
				State:             "Attica",
				Zip:               "142 42",
				IsShippingAddress: false,
			},
			&customer.CustomerRequest_Address{
				Street:            "Aristotelous 4242",
				City:              "Thessaloniki",
				State:             "Thessaloniki",
				Zip:               "542 42",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, tux)

	filter := &customer.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}
