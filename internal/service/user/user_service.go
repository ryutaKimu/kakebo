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

var dummyHash []byte

func NewUserService(pg *postgres.Postgres, userRepository repository.UserRepository) (interfaces.UserService, error) {
	if err := initDummyHash(); err != nil {
		return nil, err
	}
	return &UserServiceImpl{
		pg:             pg,
		userRepository: userRepository,
	}, nil
}

func initDummyHash() error {
	h, err := bcrypt.GenerateFromPassword([]byte("dummy-password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	dummyHash = h
	return nil
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

func (s *UserServiceImpl) Login(ctx context.Context, email string, password string) (bool, error) {
	user, err := s.userRepository.LoginUser(ctx, email)
	if err != nil {
		return false, err
	}

	var hashedPassword string
	if user == nil {
		hashedPassword = string(dummyHash)
	} else {
		hashedPassword = user.Password
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
