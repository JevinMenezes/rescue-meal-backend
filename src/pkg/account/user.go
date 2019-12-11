package account

import "context"

// User defines the properties of a user to be added
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//Role      string `json:"role"`
	//Timestamp string `json:"timestamp"`
}

// Repository - comment for export
type Repository interface {
	CreateUser(ctx context.Context, user User) error
}
