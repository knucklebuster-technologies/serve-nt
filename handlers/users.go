package handlers

import (
	"encoding/json"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("WELCOME TO USERS")
}
