package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func index(r *mux.Router) {
	r.HandleFunc("/", indexGet)
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/index.html", time.Now().String(), w)
}
