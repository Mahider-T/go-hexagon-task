package main

import (
	"encoding/gob"
	"fmt"
	"go-hexagon-task/internal/adapters/handlers"
	"go-hexagon-task/internal/adapters/storage"
	"go-hexagon-task/internal/core/service"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/justinas/alice"
)

var store *sessions.FilesystemStore

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewFilesystemStore(
		"",
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(handlers.User{})
}

func main() {

	SetEnv()

	db, err := storage.ConnectDB()

	if err != nil {
		fmt.Println("Can't connect to database")
		log.Fatal(err)
	}

	//user repo, service, handler
	userRepo := storage.NewUserRepository(db)
	userService := service.CreateUserService(userRepo)
	uh := handlers.NewUserHandler(userService)

	//task repo, service, handler
	taskRepo := storage.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	th := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", uh.Home)
	mux.HandleFunc("/user/register", uh.Register)
	mux.HandleFunc("/user/get", uh.GetUser)
	mux.HandleFunc("/user/remove", uh.Remove)
	mux.HandleFunc("/user/login", uh.Login)
	mux.HandleFunc("/user/logout", uh.Logout)

	protected := alice.New(authMiddleware)

	mux.Handle("/task/create", protected.ThenFunc(th.AddTask))
	mux.HandleFunc("/task/update", th.UpdateTask)
	mux.HandleFunc("/task/list", th.ListTasks)

	fmt.Println(`Listening on :4000`)

	err = http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
