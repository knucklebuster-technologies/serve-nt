package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qawarrior/loggy"
)

const apiVersion = "/api/v1"

func GetHandler() http.Handler {
	// rtr is the root application router
	rtr := mux.NewRouter().StrictSlash(true)
	index(rtr)
	login(rtr)
	registration(rtr)

	// Handle assest serving
	rtr.HandleFunc("/assets/css/{file}", cssGet).Methods("GET")
	rtr.HandleFunc("/assets/js/{file}", jsGet).Methods("GET")

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()
	users(api)
	// return base router rtr to the caller
	return rtr
}

func cssGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

func serveTemplate(t string, d interface{}, w io.Writer) {
	pt, err := template.ParseFiles(t)
	if err != nil {
		loggy.Error(err.Error())
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
