package main

import (
	"fmt"
	"go-hexagon-task/internal/adapters/handlers"
	"go-hexagon-task/internal/adapters/storage"
	"go-hexagon-task/internal/core/service"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	db, err := storage.ConnectDB()

	if err != nil {
		fmt.Println("Can't connect to database")
		log.Fatal(err)
	}

	//service, repo, handler

	userRepo := storage.NewUserRepository(db)
	userService := service.CreateUserService(userRepo)
	us := handlers.NewUserServiceHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user/register", us.Register)
	mux.HandleFunc("/user/get", us.GetUser)
	mux.HandleFunc("/user/remove", us.Remove)

	fmt.Println(`Listening on :4000`)

	err = http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
