package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/qawarrior/secrets"
	"github.com/qawarrior/serve-nt/model"
)

type registration struct {
	users *model.UsersCollection
}

func (h *registration) get(w http.ResponseWriter, r *http.Request) {
	p := model.PageData{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}
	serveTemplate(w, "./assets/templates/registration.html", p)
}

func (h *registration) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := &model.User{}
	err := fDecoder.Decode(u, r.PostForm)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pwd, err := secrets.HashPassword(u.Password)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Password = pwd
	u, err = h.users.Insert(u)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(u)
}
