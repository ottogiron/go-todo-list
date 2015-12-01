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
func (tasks *Tasks) Register(server server.Server) {
	server.HandleFunc("/", handleHello).Methods("GET")
}

// Name identifier for controller
func (tasks *Tasks) Name() string {
	return "TasksController"
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Tasks Controller")
}
