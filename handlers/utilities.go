package handlers

import (
	"html/template"
	"net/http"
)

// SHARED INTERNAL FUNCTIONS
func cssGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	cfg.Logger.Info.Println("Serving css -", path)
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	cfg.Logger.Info.Println("Serving js -", path)
	http.ServeFile(w, r, path)
}

func serveTemplate(w http.ResponseWriter, t string, d interface{}) {
	cfg.Logger.Info.Println("Serving template -", t)
	pt, err := template.ParseFiles(t)
	if err != nil {
		cfg.Logger.Error.Println("Failed to parse template:", t, "error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pt.Execute(w, d)
}

func sendfourOhFour(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 404)
}

func authenicated(r *http.Request) bool {
	session, _ := sessionStore.Get(r, "SNT-SESSION")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}
