package handlers

import (
	"net/http"
	"time"
)

func indexGet(w http.ResponseWriter, r *http.Request) {
	p := pagedata{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}
	serveTemplate(w, "./assets/templates/index.html", p)
}
