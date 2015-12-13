package controllers

import (
	"fmt"
	"net/http"

	"github.com/ottogiron/chapi/server"
)

// Todos handle tasks requests
type Todos struct {
	*server.BasePlugin
}

// Register registers this plugin in the server
func (t *Todos) Register(server server.Server) {
	server.HandleFunc("/todos", handleGet).Methods("GET")
}

// Name identifier for controller
func (t *Todos) Name() string {
	return "TodosController"
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		 [
			{
				"type": "todo",
				"id": 1,
				"title": "My First task",
				"isCompleted": false
			},
			{
				"type": "todo",
				"id": 2,
				"title": "My First Second task",
				"isCompleted": true
			}
		]
		`)
}
