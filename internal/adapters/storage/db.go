package storage

import (
	// "fmt"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func ConnectDB() (*Database, error) {
	DB_URL := os.Getenv("DATABASE_URL")
	fmt.Println(DB_URL)
	if DB_URL == "" {
		fmt.Println("No database URL ")
		return nil, errors.New("no database connection string provided")
	}
	dsn := flag.String("dsn", DB_URL, "Postgres data source name")
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
