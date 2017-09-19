package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/models"
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
	user := models.NewUser()
	err := fDecoder.Decode(user, r.PostForm)
	if err != nil {
		sendfourOhFour(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
