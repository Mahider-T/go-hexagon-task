package main

import (
	"fmt"
	"go-hexagon-task/internal/adapters/handlers"
	"go-hexagon-task/internal/adapters/storage"
	"go-hexagon-task/internal/core/service"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		fmt.Fprintf(w, "Not Found")
		return
	}
	w.Write([]byte("hello"))
}

func main() {

	db, err := storage.ConnectDB()

	if err != nil {
		fmt.Println("Can't connect to database")
		log.Fatal(err)
	}

	//user repo, service, handler
	userRepo := storage.NewUserRepository(db)
	userService := service.CreateUserService(userRepo)
	us := handlers.NewUserHandler(userService)

	//user repo, service, handler
	taskRepo := storage.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	th := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user/register", us.Register)
	mux.HandleFunc("/user/get", us.GetUser)
	mux.HandleFunc("/user/remove", us.Remove)
	mux.HandleFunc("/user/login", us.Login)
	mux.HandleFunc("/user/logout", us.Logout)

	protected := alice.New(authMiddleware)
	// handler := AuthMiddleware(http.HandlerFunc(th.AddTask))
	mux.Handle("/task/create", protected.ThenFunc(th.AddTask))

	mux.HandleFunc("/task/update", th.UpdateTask)
	mux.HandleFunc("/task/list", th.ListTasks)

	fmt.Println(`Listening on :4000`)

	err = http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
