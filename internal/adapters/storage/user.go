package storage

import (
	"go-hexagon-task/internal/core/domain"
)

func (rp *Database) CreateUser(usr *domain.User) {

	stmt := `INSERT INTO users (id, name, username, password, createdAt)
	VALUES(?, ?, ?,?, UTC_TIMESTAMP())`
	rp.db.Exec(stmt, usr.id, usr.name, usr.username, usr.password)
}
