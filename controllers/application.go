package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/qawarrior/serve-nt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Application represents the controller for operating on the User resource
type Application struct {
	collection *mgo.Collection
}

// NewApplication returns a controller for the User Endpoint
func NewApplication(d *mgo.Database) (*Application, error) {
	index := mgo.Index{
		Key:        []string{},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	c := d.C("applications")
	err := c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}
	return &Application{c}, nil
}

// Create adds a new Application
func (c Application) Create(w http.ResponseWriter, r *http.Request) {
	m := models.Application{}
	err := m.Decode(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	m.ID = bson.NewObjectId()
	c.collection.Insert(&m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	m.Encode(w)
}
