package service

import (
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
	"go-hexagon-task/internal/utils"

	"golang.org/x/crypto/bcrypt"
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
	hashedPass, err := utils.HashPassword(usr.Password)
	usr.Password = hashedPass

	if err != nil {
		return nil, err
	}

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

func (us UserService) Login(username string, password string) (*domain.User, error) {
	usr, err := us.userRepo.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePass(usr.Password, password)

	if err != nil {
		fmt.Println(err)
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, err
	}
	return usr, nil
}
