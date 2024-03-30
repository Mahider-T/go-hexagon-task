package storage

import (
	"fmt"
	"go-hexagon-task/internal/core/domain"
)

type UserRepository struct {
	db *Database
}

func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(usr *domain.User) (*domain.User, error) {

	stmt := `INSERT INTO users (id, name, username, password, createdAt) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)`
	_, err := ur.db.db.Exec(stmt, usr.Id, usr.Name, usr.Username, usr.Password)
	if err != nil {
		fmt.Println("Error at storage query execution")
		return nil, err
	}
	return usr, nil
}

func (us *UserRepository) GetUserById(id string) (*domain.User, error) {

	stmt := `SELECT id, name, username, password, FROM users
	WHERE id = ?`

	usr := &domain.User{}
	row := us.db.db.QueryRow(stmt, id)

	err := row.Scan(usr.Id, usr.Name, usr.Username, usr.Password)

	if err != nil {
		return nil, err
	}

	return usr, nil

}
func (us UserRepository) DeleteUser(id string) error {
	stmt := `DELETE FROM users WHERE id = ?`

	_, err := us.db.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
