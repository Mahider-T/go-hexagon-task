package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
	"net/http"
	"strconv"
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

func NewUserHandler(ser port.UserService) *UserHandler {
	return &UserHandler{
		ser,
	}
}

// Func below saves correct timestamp to db but does not display
// the time in db but rather the default value of time.Time
func (uh UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser *domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser)
	if err != nil {
		fmt.Fprintf(w, "error decoding json")
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
	fmt.Fprintf(w, "Successfully retrieved user : %v", usr)

}

func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Can't convert string to int %v", err)

	}
	usr, err := uh.ser.GetUser(id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintf(w, "No user with the given id in database")
			return
		}
		fmt.Fprintf(w, "An error occurred when trying to get the user")
		return
	}

	fmt.Fprintf(w, "Successfully retrieved user : %v", usr)
}

func (uh UserHandler) Remove(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		fmt.Fprintf(w, "Can't convert string to int: %v", err)
		return
	}

	err = uh.ser.Remove(id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintf(w, "No user with id %v", id)
			return
		}
		fmt.Fprintf(w, "Can't remove user %v", err)
		return
	}
	fmt.Fprintf(w, "Successfully removed user")
}
