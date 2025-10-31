package interfaces

import (
	"context"
	"errors"
)

var ErrUserAlreadyExists = errors.New("このメールアドレスはすでに存在しています")

type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
}
