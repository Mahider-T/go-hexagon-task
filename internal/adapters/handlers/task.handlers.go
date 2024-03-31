package handlers

import (
	"encoding/json"
	"fmt"
	"go-hexagon-task/internal/core/domain"
	"go-hexagon-task/internal/core/port"
	"net/http"
)

type TaskHandler struct {
	ser port.TaskService
}

type TaskCreateRequest struct {
	Id          int               `json: "id"`
	Title       string            `json: "title"`
	Description string            `json: "description"`
	Status      domain.TaskStatus `json: "status"`
}

func NewTaskHandler(ser port.TaskService) *TaskHandler {
	return &TaskHandler{
		ser,
	}
}
func (th TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newTask *domain.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		fmt.Fprintf(w, "Could not decode json")
		return
	}

	err = th.ser.AddTask(newTask)

	if err != nil {
		fmt.Fprintf(w, "Could not add task")
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "Task Added Successfully")
}
