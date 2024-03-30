package handlers

import (
	"encoding/json"
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
	"net/http"
)

type UserHandler struct {
	ser port.UserService
}

type UserCreateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserServiceHandler(ser port.UserService) *UserHandler {
	return &UserHandler{
		ser,
	}
}

func (uh UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var newUser *domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// w.Write([]byte("Error decoding json"))
		fmt.Fprintf(w, "error decoding json: %v", err)
		fmt.Println(err)
		return
	}

	usr, err := uh.ser.Register(newUser)
	if err != nil {
		fmt.Fprintf(w, "Error registering user : %v", err)
		fmt.Println(err)
		return
	}
	fmt.Println(usr)
	fmt.Fprintf(w, "Successfully retrieved user with name: %v", usr)

}

func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User gotten successfully!"))
}

func (uh UserHandler) Remove(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User removed successfully!"))
}
