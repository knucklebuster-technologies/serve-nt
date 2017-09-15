package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

// Controller defines a type that controls an http endpoint
type Controller interface {
	Create(http.ResponseWriter, *http.Request)
	Read(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func newIndex(k []string) mgo.Index {
	i := mgo.Index{
		Key:        k,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	return i
}

func sendfourOhFour(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(err)
}
