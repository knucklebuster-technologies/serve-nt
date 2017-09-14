package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func users(r *mux.Router) {
	r.HandleFunc("/users", getUsers)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("WELCOME TO USERS")
}
