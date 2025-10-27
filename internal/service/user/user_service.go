package service

import (
	"context"
	"errors"

	"github.com/ryutaKimu/kakebo/internal/controller/services"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) services.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, name string, email string, password string) error {
	exist, err := s.userRepository.CheckUserExists(ctx, email)

	if err != nil {
		return err
	}

	if exist {
		return errors.New("このメールアドレスはすでに存在しています")
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
	return s.userRepository.CreateUser(ctx, user)
}
