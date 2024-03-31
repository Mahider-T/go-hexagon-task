package storage

import (
	"database/sql"
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

// TODO : Show different error if user with the given id does not exist in
// the database
func (us *UserRepository) GetUserById(id int) (*domain.User, error) {

	stmt := `SELECT * FROM users WHERE id = $1`

	usr := &domain.User{}
	row := us.db.db.QueryRow(stmt, id)

	err := row.Scan(&usr.Id, &usr.Name, &usr.Username, &usr.Password, &usr.Createdat)

	if err != nil {
		return nil, err
	}

	return usr, nil

}
func (us UserRepository) DeleteUser(id int) error {

	var exists bool
	checkStmt := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"

	err := us.db.db.QueryRow(checkStmt, id).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return sql.ErrNoRows
	}

	deleteStmt := `DELETE FROM users WHERE id = $1`

	_, err = us.db.db.Exec(deleteStmt, id)

	if err != nil {
		return err
	}
	return nil
}
