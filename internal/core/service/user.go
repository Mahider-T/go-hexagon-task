package service

import (
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func CreateUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (us UserService) Register(usr domain.User) (*domain.User, error) {
	user, err := us.repo.CreateUser(usr)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us UserService) GetUser(id string) (*domain.User, error) {

	user, err := us.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us UserService) Remove(id string) error {

	err := us.repo.DeleteUser(id)

	if err != nil {
		return err
	}
	return nil
}
