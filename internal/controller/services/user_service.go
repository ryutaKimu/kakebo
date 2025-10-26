package services

import "context"

type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) error
}
