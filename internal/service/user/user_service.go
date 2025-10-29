package service

import (
	"context"

	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/user"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	pg             *postgres.Postgres
	userRepository repository.UserRepository
}

func NewUserService(pg *postgres.Postgres, userRepository repository.UserRepository) interfaces.UserService {
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
			return interfaces.ErrUserAlreadyExists
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

func (s *UserServiceImpl) Login(ctx context.Context, email string, password string) error {
	return nil
}
