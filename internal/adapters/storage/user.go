package storage

import (
	"go-hexagon-task/internal/core/domain"
)

func (rp *Database) CreateUser(usr *domain.User) (*domain.User, error) {

	stmt := `INSERT INTO users (id, name, username, password, createdAt)
	VALUES(?, ?, ?,?, UTC_TIMESTAMP())`
	_, err := rp.db.Exec(stmt, usr.Id, usr.Name, usr.Username, usr.Password)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (rp *Database) GetUserById(id string) (*domain.User, error) {

	stmt := `SELECT id, name, username, password, FROM users
	WHERE id = ?`

	us := &domain.User{}
	row := rp.db.QueryRow(stmt, id)

	err := row.Scan(us.Id, us.Name, us.Username, us.Password)

	if err != nil {
		return nil, err
	}

	return us, nil

}
func (rp *Database) DeleteUser(id string) error {
	stmt := `DELETE FROM users WHERE id = ?`

	_, err := rp.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
