package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func login(r *mux.Router) {
	r.HandleFunc("/login", loginGet).Methods("GET")
	r.HandleFunc("/login", loginPost).Methods("POST")
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/login.html", time.Now().String(), w)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	json.NewEncoder(w).Encode(r.Form)
}
