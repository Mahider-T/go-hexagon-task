package port

import (
	"go-hexagon-task/internal/core/domain"
)

type UserService interface {
	Register(usr *domain.User) (*domain.User, error)
	Login(username string, password string) (*domain.User, error)
	GetUser(id int) (*domain.User, error)
	Remove(id int) error
}

type UserRepository interface {
	CreateUser(usr *domain.User) (*domain.User, error)
	GetUserById(id int) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	DeleteUser(id int) error
}
