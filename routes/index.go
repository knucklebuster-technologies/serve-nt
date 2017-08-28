package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(router *mux.Router) {
	resource := "/"
	router.HandleFunc(resource, indexHandler).Methods("GET")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!")
}
