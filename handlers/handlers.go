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

	rtr.HandleFunc("/login", loginGet).Methods(http.MethodGet)
	rtr.HandleFunc("/login", loginPost).Methods(http.MethodPost)

	rtr.HandleFunc("/registration", registrationGet).Methods(http.MethodGet)
	rtr.HandleFunc("/registration", registrationPost).Methods(http.MethodPost)

	rtr.HandleFunc("/profile/{id}", profileidGet).Methods(http.MethodGet)

	rtr.HandleFunc("/assets/css/{file}", cssGet).Methods(http.MethodGet)
	rtr.HandleFunc("/assets/js/{file}", jsGet).Methods(http.MethodGet)

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()

	// handles for api endpoint events
	h := events{collection: models.NewEventsCollection()}
	api.HandleFunc("/events", h.post).Methods(http.MethodPost)
	api.HandleFunc("/events", h.get).Methods(http.MethodGet)

	// api handlers for users endpoint
	api.HandleFunc("/users", getUsers).Methods(http.MethodGet)

	// return base router rtr to the caller
	return rtr
}
