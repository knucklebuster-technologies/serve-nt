package routes

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

const apiVersion = "api/v1"

// Set creates router and sets up the applications routes
func Set(dbname string, db *mgo.Session) (*mux.Router, error) {
	router := mux.NewRouter().StrictSlash(true)

	// setup api routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	servent(dbname, db, apiRouter)
	servee(dbname, db, apiRouter)
	event(dbname, db, apiRouter)

	// setup page routes
	app(router)
	return router, nil
}
