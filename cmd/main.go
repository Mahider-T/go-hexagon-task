package main

import (
	"fmt"
	"go-hexagon-task/internal/adapters/handlers"
	"go-hexagon-task/internal/adapters/storage"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

type Application struct {
}

func main() {
	_, err := storage.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user/register")
	mux.HandleFunc("/user/get", handlers.GetUser)
	mux.HandleFunc("/user/remove", handlers.Remove)

	fmt.Println(`Listening on :4000`)

	err = http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
