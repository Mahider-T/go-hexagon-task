package storage

import (
	// "fmt"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func ConnectDB() (*Database, error) {
	dsn := flag.String("dsn", "postgresql://postgres:Maverick2020!@localhost:5432/go-hexagon-task", "Postgres data source name")
	db, err := sql.Open("postgres", *dsn)

	if err != nil {
		fmt.Println("Can't open DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Can't ping db")
		return nil, err
	}

	return &Database{
		db: db,
	}, nil

}
