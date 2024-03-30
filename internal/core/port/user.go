package port

import (
	"go-hexagon-task/internal/core/domain"
)

type UserService interface {
	Register(usr *domain.User) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	Remove(id string) error
}

type UserRepository interface {
	CreateUser(usr *domain.User) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	DeleteUser(id string) error
}
