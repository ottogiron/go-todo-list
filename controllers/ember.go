package controllers

import (
	"fmt"
	"net/http"

	"github.com/ottogiron/chapi/server"
)

const (
	basePath    string = "client/dist"
	defaultPath        = "client/dist/index.html"
)

type Ember struct {
	*server.BasePlugin
}

// Register registers this plugin in the server
func (e *Ember) Register(server server.Server) {
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Ember controller")
	}).Methods("GET")
}

// Name identifier for controller
func (e *Ember) Name() string {
	return "EmberController"
}
