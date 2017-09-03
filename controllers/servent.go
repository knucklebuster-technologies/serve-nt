package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qawarrior/serve-nt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Servent represents the controller for operating on the User resource
type Servent struct {
	collection *mgo.Collection
}

// NewServent returns a controller for the User Endpoint
func NewServent(d *mgo.Database) (*Servent, error) {
	c := d.C("servents")
	i := newIndex([]string{"username", "password"})
	err := c.EnsureIndex(i)
	if err != nil {
		return nil, err
	}
	return &Servent{c}, nil
}

// Create adds a new Servent
func (c Servent) Create(w http.ResponseWriter, r *http.Request) {
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

// Read returns an existing Servent
func (c Servent) Read(w http.ResponseWriter, r *http.Request) {
	m := models.Servent{}
	if err := c.collection.Find(bson.M{}).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	m.Encode(w)
}

// Update modifies an existing Servent
func (c Servent) Update(w http.ResponseWriter, r *http.Request) {
	m := models.Servent{}
	err := m.Decode(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	m.Encode(w)
}

// Delete removes an existing Servent
func (c Servent) Delete(w http.ResponseWriter, r *http.Request) {
	m := models.Servent{}
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
