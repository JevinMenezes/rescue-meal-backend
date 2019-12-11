package login

import "context"

// Endpoint describes a service as an endpoint.
type Endpoint func(ctx context.Context, req interface{}) (resp interface{}, err error)

// MakeLoginEndpoint describes the login microservice endpoint
func MakeLoginEndpoint(s Service) Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		request := req.(Request)
		s, err := s.Login(ctx, request.user)
		return Response{S: s, Err: err}, nil
	}
}

// Request describes data fields of the login service request
type Request struct {
	user User
}

// Response describes data fields of the login service response
type Response struct {
	S   string
	Err error
}
