package server

import (
	"encoding/json"
	"net/http"

	"github.com/SevvyP/todo_web_v1/internal/util"
)

// GetTasks handles the GET /tasks route. It will fetch all tasks for
// the current user formatted as a json document with each task nested
// under its parent, and return an error if no user is present in the request.
// If an "id" query parameter is present, it will return a json document
// with the task and any subtasks nested under it, or 404 if the task is not found.
func (r *Resolver) GetTasks(w http.ResponseWriter, req *http.Request) {
	userID := "get this from the session or something idk"

	if userID == "" {
		http.Error(w, "No user found", http.StatusBadRequest)
		return
	}

	id := req.URL.Query().Get("id")

	// Check for "id" query parameter
	if id != "" {
		r.GetTasksByID(w, req, userID, id)
		return
	}

	tasks, err := r.TaskService.GetTasks(userID)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	out, err := util.FormatTaskJson(tasks)
	if err != nil {
		http.Error(w, "Failed to format tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

// GetTasksByID handles the GET /tasks?id=<id> route. It will fetch the task
// with the given id and return a json document with the task and any subtasks
// nested under it, or 404 if the task is not found.
func (r *Resolver) GetTasksByID(w http.ResponseWriter, req *http.Request, userID string, id string) {
	tasks, err := r.TaskService.GetTasksByID(userID, id)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	if tasks == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	out, err := util.FormatTaskJson(tasks)
	if err != nil {
		http.Error(w, "Failed to format tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(out)
}

func (r *Resolver) CreateTask(w http.ResponseWriter, req *http.Request) {

}

func (r *Resolver) UpdateTask(w http.ResponseWriter, req *http.Request) {

}

func (r *Resolver) DeleteTask(w http.ResponseWriter, req *http.Request) {

}
