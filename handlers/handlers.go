package handlers

import (
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

	// api is a subrouter of the root router rtr
	api := rtr.PathPrefix(apiVersion).Subrouter()
	users(api)
	// return base router rtr to the caller
	return rtr
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
