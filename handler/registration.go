package handler

import (
	"encoding/json"
	"html/template"
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
	tpl, err := template.ParseFiles("./assets/templates/_layout.html", "./assets/templates/registration.html")
	if err != nil {
		return
	}
	tpl.ExecuteTemplate(w, "_layout", p)
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
