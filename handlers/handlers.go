package handlers

import (
	"html/template"
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
	api.HandleFunc("/users", getUsers).Methods(http.MethodGet)

	// set config for package
	cfg = c

	// return base router rtr to the caller
	return rtr
}

// SHARED INTERNAL FUNCTIONS

func cssGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	cfg.Logger.Info.Println("Serving css -", path)
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	cfg.Logger.Info.Println("Serving js -", path)
	http.ServeFile(w, r, path)
}

func serveTemplate(w http.ResponseWriter, t string, d interface{}) {
	cfg.Logger.Info.Println("Serving template -", t)
	pt, err := template.ParseFiles(t)
	if err != nil {
		cfg.Logger.Error.Println("Failed to parse template:", t, "error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pt.Execute(w, d)
}

func sendfourOhFour(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 404)
}

func authenicated(r *http.Request) bool {
	session, _ := sessionStore.Get(r, "SNT-SESSION")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}
