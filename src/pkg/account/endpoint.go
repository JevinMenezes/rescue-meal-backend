package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints - comment for export
type Endpoints struct {
	CreateUser endpoint.Endpoint
}

// MakeEndpoints - comment for export
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoints(s),
	}
}

func makeCreateUserEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}
