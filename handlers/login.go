package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func login(r *mux.Router) {
	r.HandleFunc("/login", loginGet)
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/login.html", time.Now().String(), w)
}
