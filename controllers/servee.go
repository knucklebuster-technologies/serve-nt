package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qawarrior/serve-nt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Servee represents the controller for operating on the User resource
type Servee struct {
	collection *mgo.Collection
}

// NewServee returns a controller for the User Endpoint
func NewServee(d *mgo.Database) (*Servee, error) {
	c := d.C("servees")
	i := newIndex([]string{"username", "password"})
	err := c.EnsureIndex(i)
	if err != nil {
		return nil, err
	}
	return &Servee{c}, nil
}

// Create adds a new Servee
func (c Servee) Create(w http.ResponseWriter, r *http.Request) {
	m := models.Servee{}
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

// Read returns an existing Servee
func (c Servee) Read(w http.ResponseWriter, r *http.Request) {
	m := models.Servee{}
	if err := c.collection.Find(bson.M{}).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	m.Encode(w)
}

// Update modifies an existing Servee
func (c Servee) Update(w http.ResponseWriter, r *http.Request) {
	m := models.Servee{}
	err := m.Decode(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	m.Encode(w)
}

// Delete removes an existing Servee
func (c Servee) Delete(w http.ResponseWriter, r *http.Request) {
	m := models.Servee{}
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
