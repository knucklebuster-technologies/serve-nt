package handlers

import (
	"net/http"

	"github.com/qawarrior/serve-nt/configuration"
	"github.com/qawarrior/serve-nt/models"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

const apiVersion = "/api/v1"

var (
	cfg          *configuration.Config
	fDecoder     = schema.NewDecoder()
	storeKey     = "Serving Neighbors Transportation"
	sessionStore = sessions.NewCookieStore([]byte(storeKey))
)

// New return the router for the application
func New(c *configuration.Config) http.Handler {
	// setup the models package for use with the handlers
	models.Init(c)

	// set config for package
	cfg = c

	// rtr is the root application router
	rtr := mux.NewRouter().StrictSlash(true)
	rtr.HandleFunc("/", indexGet)

	ucol := models.NewUsersCollection()
	l := login{users: ucol}
	rtr.HandleFunc("/login", l.get).Methods(http.MethodGet)
	rtr.HandleFunc("/login", l.post).Methods(http.MethodPost)

	r := registration{collection: ucol}
	rtr.HandleFunc("/registration", r.get).Methods(http.MethodGet)
	rtr.HandleFunc("/registration", r.post).Methods(http.MethodPost)

	p := profile{users: ucol, events: models.NewEventsCollection()}
	rtr.HandleFunc("/profile/{id}", p.get).Methods(http.MethodGet)

	rtr.HandleFunc("/assets/css/{file}", cssGet).Methods(http.MethodGet)
	rtr.HandleFunc("/assets/js/{file}", jsGet).Methods(http.MethodGet)

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()

	// handles for api endpoint events
	e := events{collection: models.NewEventsCollection()}
	api.HandleFunc("/events", e.post).Methods(http.MethodPost)
	api.HandleFunc("/events", e.get).Methods(http.MethodGet)

	// api handlers for users endpoint
	api.HandleFunc("/users", getUsers).Methods(http.MethodGet)

	// return base router rtr to the caller
	return rtr
}
