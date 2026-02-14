package accounts

import "context"

type Service interface {
	RegisterUser(ctx context.Context, email, password, name string) (User, error)
	Authenticate(ctx context.Context, email, password string) (User, error)
	GetUser(ctx context.Context, id string) (User, error)
	ListUsers(ctx context.Context, limit int, cursor string) ([]User, string, error)
	DisableUser(ctx context.Context, id string) error
}
