package main

import (
	"go-hexagon-task/internal/adapters/storage"
	"log"
)

func main() {
	_, err := storage.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
}
