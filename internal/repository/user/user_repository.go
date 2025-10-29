package repository

import (
	"context"

	"github.com/ryutaKimu/kakebo/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	CheckUserExists(ctx context.Context, email string) (bool, error)
	LoginUser(ctx context.Context, email string) (*model.User, error)
}
