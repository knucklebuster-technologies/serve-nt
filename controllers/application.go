package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/qawarrior/serve-nt/loggy"
)

// Application represents the controller for operating on the application pages
type Application struct {
}

// NewApplication returns a controller for the Application pages
func NewApplication() (*Application, error) {
	loggy.Info("NEW Application CONTROLLER BEING CREATED")
	return &Application{}, nil
}

// index handlers

// IndexGet handles serving the index.html template
func (c Application) IndexGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.IndexGet CALLED")
	serveTemplate("./assets/templates/index.html", time.Now().String(), w)
}

// login handlers

// LoginGet handles serving the login.html template
func (c Application) LoginGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.loginGet CALLED")
	serveTemplate("./assets/templates/login.html", time.Now().String(), w)
}

// LoginPost handles the submitted form data
func (c Application) LoginPost(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.loginPost CALLED")
	r.ParseForm()
	loggy.Info(r.Form)
	json.NewEncoder(w).Encode(r.Form)
}

// registration handlers

// RegisterGet handles serving the register.html template
func (c Application) RegisterGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.RegisterGet CALLED")
	serveTemplate("./assets/templates/register.html", time.Now().String(), w)
}

// RegisterPost handles the submitted form data
func (c Application) RegisterPost(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.RegisterPost CALLED")
	r.ParseForm()
	loggy.Info(r.Form)
	json.NewEncoder(w).Encode(r.Form)
}

// asset handlers

// CSSGet handles serving style sheets
func (c Application) CSSGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.CSSGet CALLED")
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

// JSGet handles serving javascripts
func (c Application) JSGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Application.JSGet CALLED")
	path := "." + r.URL.Path
	http.ServeFile(w, r, path)
}

// Serve pages
func serveTemplate(t string, d interface{}, w io.Writer) {
	pt, err := template.ParseFiles(t)
	if err != nil {
		loggy.Error(err.Error())
		fmt.Fprintln(w, err)
		return
	}
	pt.Execute(w, d)
}
