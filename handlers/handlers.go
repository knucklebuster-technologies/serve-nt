package handlers

import (
	"html/template"
	"net/http"

	"github.com/qawarrior/serve-nt/configuration"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

const apiVersion = "/api/v1"

var (
	fDecoder     = schema.NewDecoder()
	storeKey     = "Serving Neighbors Transportation"
	sessionStore = sessions.NewCookieStore([]byte(storeKey))
)

// GetHandler returns the route handler for this application
func GetHandler() http.Handler {
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

	// return base router rtr to the caller
	return rtr
}

// SHARED INTERNAL FUNCTIONS

func cssGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	configuration.Linfo.Println("Serving css -", path)
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	configuration.Linfo.Println("Serving js -", path)
	http.ServeFile(w, r, path)
}

func serveTemplate(t string, d interface{}, w http.ResponseWriter) {
	configuration.Linfo.Println("Serving template -", t)
	pt, err := template.ParseFiles(t)
	if err != nil {
		configuration.Lerror.Println("Failed to parse template:", t, "error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pt.Execute(w, d)
}

func sendfourOhFour(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 404)
}
