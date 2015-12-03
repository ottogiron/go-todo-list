package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/ottogiron/chapi/server"
	"github.com/wunderlist/moxy"
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

	hosts := []string{"localhost:4200"}
	proxy := moxy.NewReverseProxy(hosts, nil)
	server.HandleFunc("/ember-cli-live-reload.js", proxy.ServeHTTP).Methods("GET")

	server.HandleFunc("/{params:.*}", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" {
			writeIndex(w)

		} else {
			resourcePath := path.Join(basePath, r.URL.Path)
			file, err := os.Open(resourcePath)
			defer file.Close()

			if err != nil {
				writeIndex(w)
				return
			}
			fi, _ := file.Stat()
			isFile := fi.Mode().IsRegular()

			if isFile {
				writeFile(w, resourcePath)
			}
		}
	}).Methods("GET")

}

func writeIndex(w http.ResponseWriter) {
	writeFile(w, defaultPath)
}

func writeFile(w http.ResponseWriter, path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(file)
}

// Name identifier for controller
func (e *Ember) Name() string {
	return "EmberController"
}
