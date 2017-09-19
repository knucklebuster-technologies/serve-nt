package routes

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

const apiVersion = "/api/v1"

var rtr = mux.NewRouter().StrictSlash(true)
var api = rtr.PathPrefix(apiVersion).Subrouter()

// Set creates router and sets up the applications routes
func Set(dbname string, db *mgo.Session) (*mux.Router, error) {
	servent(dbname, db, api)
	servee(dbname, db, api)
	event(dbname, db, api)

	return rtr, nil
}
