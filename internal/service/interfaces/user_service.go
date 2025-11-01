package interfaces

import (
	"context"
	"errors"

	"github.com/ryutaKimu/kakebo/internal/model"
)

var ErrUserAlreadyExists = errors.New("このメールアドレスはすでに存在しています")

type UserService interface {
	CreateUser(ctx context.Context, name string, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
	GetProfile(ctx context.Context, id int) (*model.User, error)
}
