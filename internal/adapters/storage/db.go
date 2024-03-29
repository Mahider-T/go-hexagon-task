package storage

import (
	// "fmt"
	"database/sql"
	"flag"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func ConnectDB() (*Database, error) {
	dsn := flag.String("dsn", "postgresql://postgres:Maverick2020!@localhost:5432/go-hexagon-task", "Postgres data source name")
	db, err := sql.Open("postgres", *dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil

}
