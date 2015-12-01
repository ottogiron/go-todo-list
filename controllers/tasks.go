package controllers

import (
	"fmt"
	"net/http"

	"github.com/ottogiron/chapi/server"
)

// Tasks handle tasks requests
type Tasks struct {
	*server.BasePlugin
}

// Register registers this plugin in the server
func (t *Tasks) Register(server server.Server) {
	server.HandleFunc("/tasks", handleGet).Methods("GET")
}

// Name identifier for controller
func (t *Tasks) Name() string {
	return "TasksController"
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Tasks Controller")
}
