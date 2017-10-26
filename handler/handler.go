package handler

import (
	"net/http"

	"github.com/qawarrior/serve-nt/configuration"
	"github.com/qawarrior/serve-nt/model"

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
	// setup the model package for use with the handlers
	model.Init(c)

	// set config for package
	cfg = c

	// rtr is the root application router
	rtr := mux.NewRouter().StrictSlash(true)

	i := index{}
	rtr.HandleFunc("/", i.get)

	ucol := model.NewUsersCollection()
	ecol := model.NewEventsCollection()
	l := login{users: ucol}
	rtr.HandleFunc("/login", l.get).Methods(http.MethodGet)
	rtr.HandleFunc("/login", l.post).Methods(http.MethodPost)

	r := registration{users: ucol}
	rtr.HandleFunc("/registration", r.get).Methods(http.MethodGet)
	rtr.HandleFunc("/registration", r.post).Methods(http.MethodPost)

	p := profile{users: ucol, events: ecol}
	rtr.HandleFunc("/profile/{id}", p.get).Methods(http.MethodGet)

	e := event{events: ecol}
	rtr.HandleFunc("/event", e.get).Methods(http.MethodGet)
	rtr.HandleFunc("/event", e.post).Methods(http.MethodPost)

	a := assets{}
	rtr.HandleFunc("/assets/css/{file}", a.css).Methods(http.MethodGet)
	rtr.HandleFunc("/assets/js/{file}", a.js).Methods(http.MethodGet)

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()

	// handles for api endpoint events
	es := events{collection: ecol}
	api.HandleFunc("/events", es.post).Methods(http.MethodPost)
	api.HandleFunc("/events", es.get).Methods(http.MethodGet)

	// api handlers for users endpoint
	api.HandleFunc("/users", getUsers).Methods(http.MethodGet)

	// return base router rtr to the caller
	return rtr
}
