package account

import "context"

// Service - comment for export
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
}
