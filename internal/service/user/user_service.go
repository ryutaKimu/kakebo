package service

import (
	"context"

	"github.com/ryutaKimu/kakebo/internal/controller/services"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	pg             *postgres.Postgres
	userRepository repository.UserRepository
}

func NewUserService(pg *postgres.Postgres, userRepository repository.UserRepository) services.UserService {
	return &UserServiceImpl{
		pg:             pg,
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, name string, email string, password string) error {
	return s.pg.Transaction(ctx, func(txCtx context.Context) error {
		exist, err := s.userRepository.CheckUserExists(txCtx, email)
		if err != nil {
			return err
		}
		if exist {
			return services.ErrUserAlreadyExists
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user := &model.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		return s.userRepository.CreateUser(txCtx, user)
	})
}
