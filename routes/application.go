package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/loggy"
)

var queryValue = time.Now().String()

func application(router *mux.Router) {
	// index routes
	router.HandleFunc("/", indexGet).Methods("GET")

	// login routes
	router.HandleFunc("/login", loginGet).Methods("GET")
	router.HandleFunc("/login", loginPost).Methods("POST")

	// registration routes
	router.HandleFunc("/registration", registrationGet).Methods("GET")
	router.HandleFunc("/registration", registrationPost).Methods("POST")

	// asset routes
	router.HandleFunc("/assets/css/{file}", cssGet).Methods("GET")
	router.HandleFunc("/assets/js/{file}", jsGet).Methods("GET")
}

// index handlers
func indexGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Routes.indexGet CALLED")
	serveTemplate("./assets/templates/index.html", queryValue, w)
}

// login handlers
func loginGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Routes.loginGet CALLED")
	serveTemplate("./assets/templates/login.html", queryValue, w)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Routes.loginPost CALLED")
	r.ParseForm()
	loggy.Info(r.Form)
}

// registration handlers
func registrationGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Routes.registrationGet CALLED")
	serveTemplate("./assets/templates/register.html", queryValue, w)
}

func registrationPost(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER Routes.registrationPost CALLED")
	r.ParseForm()
	loggy.Info(r.Form)
}

// asset handlers
func cssGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER routes.cssGet CALLED")
	path := "." + r.URL.Path
	loggy.Info("CSS PATH REQUESTED: " + path)
	http.ServeFile(w, r, path)
}

func jsGet(w http.ResponseWriter, r *http.Request) {
	loggy.Info("HANDLER routes.jsGet CALLED")
	path := "." + r.URL.Path
	loggy.Info("JS PATH REQUESTED: " + path)
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
