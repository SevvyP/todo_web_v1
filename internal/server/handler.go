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

	userID := req.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "no user_id in request params", http.StatusBadRequest)
		return
	}

	tasks, err := r.TaskService.GetTasks(userID)
	if err != nil {
		http.Error(w, "Failed to get tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if tasks == nil {
		http.Error(w, "No tasks found", http.StatusNotFound)
		return
	}
	out, err := util.FormatTaskJson(*tasks)
	if err != nil {
		http.Error(w, "Failed to format tasks", http.StatusInternalServerError)
		return
	}
	if out == nil {
		http.Error(w, "No tasks found", http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	json.NewEncoder(w).Encode(out)
}

func (r *Resolver) CreateTask(w http.ResponseWriter, req *http.Request) {

}

func (r *Resolver) UpdateTask(w http.ResponseWriter, req *http.Request) {

}

func (r *Resolver) DeleteTask(w http.ResponseWriter, req *http.Request) {

}
