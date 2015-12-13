package models

import "gopkg.in/mgo.v2/bson"

//Todo todo model
type Todo struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `json:"title"`
	IsCompleted bool          `json:"isCompleted"`
	MType       string        `json:"type"`
}

//Todos todo lists
type Todos []Todo

//NewTodo initializes a new todo
func NewTodo() *Todo {
	id := bson.NewObjectId()
	return &Todo{ID: id, MType: "todo"}
}
