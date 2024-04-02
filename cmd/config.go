package main

import (
	"os"
)

func SetEnv() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:Maverick2020!@localhost:5432/go-hexagon-task")
}
