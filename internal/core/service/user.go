package service

import (
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
)

type UserService struct {
	userRepo port.UserRepository
}

func CreateUserService(port port.UserRepository) *UserService {
	return &UserService{
		userRepo: port,
	}
}

func (us UserService) Register(usr *domain.User) (*domain.User, error) {
	user, err := us.userRepo.CreateUser(usr)

	if err != nil {
		fmt.Println("Error at user service")
		return nil, err
	}

	return user, nil
}

func (us UserService) GetUser(id int) (*domain.User, error) {

	user, err := us.userRepo.GetUserById(id)

	if err != nil {
		fmt.Println("Error at service")
		return nil, err
	}

	return user, nil
}

func (us UserService) Remove(id int) error {

	err := us.userRepo.DeleteUser(id)

	if err != nil {
		return err
	}
	return nil
}
