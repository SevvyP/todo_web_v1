package server

import (
	"net/http"

	"github.com/SevvyP/todo_web_v1/internal/middleware"
	"github.com/SevvyP/todo_web_v1/internal/service"
)

// Resolver is the main server struct that holds the HTTP server and the database.
type Resolver struct {
	Server      *http.Server
	TaskService service.TaskService
}

type Config struct {
	AuthConfig *middleware.AuthConfig `json:"auth"`
	TaskConfig *service.TaskConfig    `json:"tasks"`
}

// NewResolver creates a new Resolver with a new HTTP server and database.
// It also sets up the HTTP routes for the server.
func NewResolver(config *Config) *Resolver {

	if config.AuthConfig == nil {
		panic("AuthConfig is required")
	}

	if config.TaskConfig == nil {
		panic("TaskConfig is required")
	}

	tokenService := service.NewTokenService(config.AuthConfig)

	mux := http.NewServeMux()
	resolver := &Resolver{
		Server: &http.Server{
			Addr:    ":8081",
			Handler: mux,
		},
		TaskService: service.NewTaskService(*config.TaskConfig, tokenService),
	}

	// Wrap the handler with the authentication middleware
	mux.Handle("/tasks", middleware.EnsureValidToken(config.AuthConfig)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodOptions:
			// validate the preflight request
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			resolver.GetTasks(w, r)
		case http.MethodPost:
			resolver.CreateTask(w, r)
		case http.MethodPut:
			resolver.UpdateTask(w, r)
		case http.MethodDelete:
			resolver.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	return resolver
}

// Resolve starts the HTTP server and listens for incoming requests.
func (r *Resolver) Resolve() error {
	return r.Server.ListenAndServe()
}
