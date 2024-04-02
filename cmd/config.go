package main

import (
	"os"
)

func SetEnv() {
	os.Setenv("DATABASE_URL", "postgresql://postgres:pass@localhost:5432/go-hexagon-task")
}
