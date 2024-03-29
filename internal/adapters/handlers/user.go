package handlers

import (
	"encoding/json"
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
	"net/http"
)

type UserService struct {
	ser port.UserService
}

type UserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (us UserService) Register(w http.ResponseWriter, r *http.Request) {

	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// w.Write([]byte("Error decoding json"))
		fmt.Fprintf(w, "error decoding json: %v", err)
		fmt.Println(err)
		return
	}

	usr, err := us.ser.Register(newUser)
	if err != nil {
		fmt.Fprintf(w, "Error registering user : %v", err)
		fmt.Println(err)
		return
	}
	fmt.Println(usr)
	fmt.Fprintf(w, "Successfully retrieved user with name: %v", usr)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User gotten successfully!"))
}

func Remove(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User removed successfully!"))
}
