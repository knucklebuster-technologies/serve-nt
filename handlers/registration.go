package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func registration(r *mux.Router) {
	r.HandleFunc("/registration", registrationGet)
}

func registrationGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/registration.html", time.Now().String(), w)
}
