package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/qawarrior/secrets"
	"github.com/qawarrior/serve-nt/models"
)

func registration(r *mux.Router) {
	r.HandleFunc("/registration", registrationGet).Methods("GET")
	r.HandleFunc("/registration", registrationPost).Methods("POST")
}

func registrationGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/registration.html", time.Now().String(), w)
}

func registrationPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := r.FormValue("password")
	h, err := secrets.HashPassword(p)
	if err != nil {
		sendfourOhFour(w, err)
	}

	m := models.Servee{}
	m.Email = r.FormValue("email")
	m.Password = h
	m.Firstname = r.FormValue("firstname")
	m.Lastname = r.FormValue("lastname")
	mb, err := json.Marshal(m)
	if err != nil {
		sendfourOhFour(w, err)
	}
	resp, err := http.DefaultClient.Post("http://127.0.0.1:8001/api/v1/servees", "application/json", bytes.NewBuffer(mb))
	if err != nil {
		sendfourOhFour(w, err)
	}
	m.Decode(resp.Body)
	defer resp.Body.Close()
	m.Encode(w)
}
