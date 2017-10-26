package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/qawarrior/serve-nt/model"
)

type event struct {
	events *model.EventsCollection
}

func (h *event) get(w http.ResponseWriter, r *http.Request) {
	if authenicated(r) != true {
		http.Error(w, "Not Authenicated", http.StatusForbidden)
		return
	}
	p := model.PageData{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}

	tpl, err := template.ParseFiles("./assets/templates/_layout.html", "./assets/templates/event.html")
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "_layout", p)
}

func (h *event) post(w http.ResponseWriter, r *http.Request) {
	if authenicated(r) != true {
		cfg.Logger.Error.Println("NOT AUTHENICATED")
		http.Error(w, "NOT AUTHENICATED", http.StatusForbidden)
		return
	}
	r.ParseForm()
	e := &model.Event{}
	err := fDecoder.Decode(e, r.PostForm)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e.ServeeID = getid(r)
	e, err = h.events.Insert(e)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(e)
}
