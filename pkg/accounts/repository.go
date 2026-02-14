package accounts

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, u User) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListUsers(ctx context.Context, limit int, cursor string) ([]User, string, error)
	UpdateUser(ctx context.Context, u User) (User, error)
	DisableUser(ctx context.Context, id string) error
}
