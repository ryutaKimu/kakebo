package service

import repository "github.com/ryutaKimu/kakebo/internal/repository/user"

type TopServiceImpl struct {
	repo repository.UserRepository
}

func NewTopService(userRepository repository.UserRepository) *TopServiceImpl {
	return &TopServiceImpl{repo: userRepository}
}
