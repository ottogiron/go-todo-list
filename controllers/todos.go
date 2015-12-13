package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ottogiron/chapi/server"
	"github.com/ottogiron/gotodo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Todos handle tasks requests
type Todos struct {
	*server.BasePlugin
}

// Register registers this plugin in the server
func (t *Todos) Register(server server.Server) {
	server.HandleFunc("/todos", handleGet).Methods("GET")
	server.HandleFunc("/todos", handlePOST).Methods("POST")
}

// Name identifier for controller
func (t *Todos) Name() string {
	return "TodosController"
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	session, err := createMongoSession()
	defer session.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	collection := session.DB("todos").C("todos")
	result := models.Todos{}
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	todo := models.NewTodo()
	if err := json.Unmarshal(body, todo); err != nil {
		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	session, err := createMongoSession()
	defer session.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collection := session.DB("todos").C("todos")
	collection.Insert(todo)
	json.NewEncoder(w).Encode(todo)

}

func createMongoSession() (*mgo.Session, error) {
	return mgo.Dial("192.168.99.100")
}
