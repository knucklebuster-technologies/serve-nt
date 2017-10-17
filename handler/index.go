package handler

import (
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
	serveTemplate(w, "./assets/templates/index.html", p)
}
