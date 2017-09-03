package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/qawarrior/serve-nt/loggy"
)

// Event represents the controller for operating on the User resource
type Event struct {
	collection *mgo.Collection
}

// NewEvent returns a controller for the User Endpoint
func NewEvent(d *mgo.Database) (*Event, error) {
	loggy.Info("NEW EVENT CONTROLLER BEING CREATED")
	c := d.C("events")
	i := newIndex([]string{"title", "servee_id"})
	err := c.EnsureIndex(i)
	if err != nil {
		loggy.Error(err)
		return nil, err
	}
	return &Event{c}, nil
}

// Create adds a new Event
func (c Event) Create(w http.ResponseWriter, r *http.Request) {
	loggy.Info("Event.Create HANDLER CALLED")
	m := models.Event{}
	err := m.Decode(r.Body)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	m.ID = bson.NewObjectId()
	if m.ServeeID == "" {
		m.ServeeID = "000000000000000000000000"
	}
	if m.ServentID == "" {
		m.ServentID = "000000000000000000000000"
	}

	err = c.collection.Insert(&m)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	m.Encode(w)
}

// Read returns all existing Events
func (c Event) Read(w http.ResponseWriter, r *http.Request) {
	loggy.Info("Event.Read HANDLER CALLED")
	m := []models.Event{}
	err := c.collection.Find(bson.M{}).All(&m)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(m)
}

// ReadTitle returns the Events with the specified Tile
func (c Event) ReadTitle(w http.ResponseWriter, r *http.Request) {
	loggy.Info("Event.ReadTitle HANDLER CALLED")
	m := []models.Event{}

	vars := mux.Vars(r)
	title := vars["title"]

	err := c.collection.Find(bson.M{"title": title}).All(&m)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(m)
}

// Update modifies an existing Event
func (c Event) Update(w http.ResponseWriter, r *http.Request) {
	loggy.Info("Event.Update HANDLER CALLED")
	m := models.Event{}
	err := m.Decode(r.Body)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	m.Encode(w)
}

// Delete removes an existing Event
func (c Event) Delete(w http.ResponseWriter, r *http.Request) {
	loggy.Info("Event.Delete HANDLER CALLED")
	m := models.Event{}
	err := m.Decode(r.Body)
	if err != nil {
		loggy.Error(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	ci, err := c.collection.RemoveAll(bson.M{})
	if err != nil {
		loggy.Error(err)
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
