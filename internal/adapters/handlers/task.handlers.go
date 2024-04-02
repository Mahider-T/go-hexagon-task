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

type TaskHandler struct {
	ser port.TaskService
}

type TaskCreateRequest struct {
	Title       string            `json: "title"`
	Description string            `json: "description"`
	Status      domain.TaskStatus `json: "status"`
}

type TaskUpdateRequest struct {
	Status domain.TaskStatus `json:"status"`
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

func (th TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "Could not convert id to string")
		return
	}

	var newStatusInput *TaskUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&newStatusInput)

	if err != nil {
		fmt.Printf("err : %v", err)
		fmt.Fprintf(w, "Could not decode json")
		return
	}
	err = th.ser.UpdateTask(id, &newStatusInput.Status)

	if err != nil {
		fmt.Fprintf(w, "Could not update task status %v", err)
		return
	}

	fmt.Fprintf(w, "Task updated successfully")
}

func (th TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := th.ser.ListTask()

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintf(w, "No tasks yet")
			return
		}
		fmt.Fprintf(w, "Could not get tasks")
		return
	}

	for id, task := range tasks {
		fmt.Fprintf(w, "%v : %v\n", id+1, task)
	}

	fmt.Fprintf(w, "These are all the tasks at hand")

}
