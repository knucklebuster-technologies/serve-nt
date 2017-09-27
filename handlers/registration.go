package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/qawarrior/secrets"
	"github.com/qawarrior/serve-nt/models"
)

func registrationGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/registration.html", tempdata{Timestamp: time.Now(), AppName: "SERVE-NT"}, w)
}

func registrationPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	m := models.NewUser()
	err := fDecoder.Decode(m, r.PostForm)
	if err != nil {
		sendfourOhFour(w, err)
		return
	}

	h, err := secrets.HashPassword(m.Password)
	if err != nil {
		sendfourOhFour(w, err)
		return
	}
	m.Password = h
	err = m.Insert()
	if err != nil {
		sendfourOhFour(w, err)
		return
	}
	json.NewEncoder(w).Encode(m)
}
