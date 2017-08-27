package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qawarrior/serve-nt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Event represents the controller for operating on the User resource
type Event struct {
	collection *mgo.Collection
}

// NewEvent returns a controller for the User Endpoint
func NewEvent(d *mgo.Database) (*Event, error) {
	index := mgo.Index{
		Key:        []string{"_id", "title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	c := d.C("events")
	err := c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}
	return &Event{c}, nil
}

// Create adds a new Event
func (c Event) Create(w http.ResponseWriter, r *http.Request) {
	m := models.Event{}
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

// Read returns an existing Event
func (c Event) Read(w http.ResponseWriter, r *http.Request) {
	m := models.Event{}
	if err := c.collection.Find(bson.M{}).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	m.Encode(w)
}

// Update modifies an existing Event
func (c Event) Update(w http.ResponseWriter, r *http.Request) {
	m := models.Event{}
	err := m.Decode(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	m.Encode(w)
}

// Delete removes an existing Event
func (c Event) Delete(w http.ResponseWriter, r *http.Request) {
	m := models.Event{}
	err := m.Decode(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	ci, err := c.collection.RemoveAll(bson.M{})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Marshal provided interface into JSON structure
	cij, _ := json.Marshal(ci)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cij)
}
