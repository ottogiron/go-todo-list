package controllers

import (
	"fmt"
	"net/http"

	"github.com/ottogiron/chapi/server"
)

// Tasks handle tasks requests
type Todos struct {
	*server.BasePlugin
}

// Register registers this plugin in the server
func (t *Todos) Register(server server.Server) {
	server.HandleFunc("/todos", handleGet).Methods("GET")
}

// Name identifier for controller
func (t *Todos) Name() string {
	return "TasksController"
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		[
				{
					"title": "My First task",
					"isCompleted": false
				}
			]
		`)
}
