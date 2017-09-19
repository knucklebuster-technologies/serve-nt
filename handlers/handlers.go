package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/qawarrior/serve-nt/configuration"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

const apiVersion = "/api/v1"

var fDecoder = schema.NewDecoder()

// GetHandler returns the route handler for this application
func GetHandler() http.Handler {
	configuration.Linfo.Println("Creating mux router")

	// rtr is the root application router
	rtr := mux.NewRouter().StrictSlash(true)
	index(rtr)
	login(rtr)
	registration(rtr)

	// Handle assest serving
	rtr.HandleFunc("/assets/css/{file}", cssGet).Methods("GET")
	rtr.HandleFunc("/assets/js/{file}", jsGet).Methods("GET")
	configuration.Linfo.Println("Application handlers applied")

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()
	users(api)
	configuration.Linfo.Println("API handlers applied")

	configuration.Linfo.Println("Returning configured handler")
	// return base router rtr to the caller
	return rtr
}

func cssGet(w http.ResponseWriter, r *http.Request) {
	configuration.Linfo.Println("CSS requested")
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	configuration.Linfo.Println("Javascript requested")
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

func serveTemplate(t string, d interface{}, w io.Writer) {
	pt, err := template.ParseFiles(t)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	pt.Execute(w, d)
}

func sendfourOhFour(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(err)
}
