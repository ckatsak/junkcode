package customer

import (
	"context"
	"strings"
)

type Server struct {
	savedCustomers []*CustomerRequest
}

func (s *Server) CreateCustomer(ctx context.Context, in *CustomerRequest) (*CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *Server) GetCustomers(filter *CustomerFilter, stream Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}
