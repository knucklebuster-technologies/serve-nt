package handlers

import (
	"net/http"
	"time"
)

func indexGet(w http.ResponseWriter, r *http.Request) {
	serveTemplate("./assets/templates/index.html", tempdata{Timestamp: time.Now(), AppName: "SERVE-NT"}, w)
}
