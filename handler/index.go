package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/qawarrior/serve-nt/model"
)

type index struct{}

func (h *index) get(w http.ResponseWriter, r *http.Request) {
	p := model.PageData{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}
	tpl, err := template.ParseFiles("./assets/templates/_layout.html", "./assets/templates/index.html")
	if err != nil {
		return
	}
	tpl.ExecuteTemplate(w, "_layout", p)
}
