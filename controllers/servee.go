package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/qawarrior/loggy"
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
	loggy.Info("CREATING Servee CONTROLLER")
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
	loggy.Info("HANDLER Servee.Create CALLED")
	m := models.Servee{}
	err := m.Decode(r.Body)
	if err != nil {
		sendfourOhFour(w, err)
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
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			sendfourOhFour(w, err)
			return
		}
		u := r.Form.Get("username")
		p := r.Form.Get("password")
		err = c.collection.Find(bson.M{"username": u, "password": p}).One(&m)
		if err != nil {
			sendfourOhFour(w, err)
			return
		}
	}

	v := mux.Vars(r)
	id := v["id"]
	err := c.collection.Find(bson.M{"_id": id}).One(&m)
	if err != nil {
		sendfourOhFour(w, err)
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
		sendfourOhFour(w, err)
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
