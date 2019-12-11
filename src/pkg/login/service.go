package login

import "context"

// Service describes a login microservice
type Service interface {
	Login(ctx context.Context, user User) (string, error)
}

type service struct{}

func (s service) Login(_ context.Context, user User) (string, error) {
	return "success", nil
}
