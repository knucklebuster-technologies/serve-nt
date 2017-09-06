package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/controllers"
)

func application(router *mux.Router) {
	c := controllers.Application{}
	// index routes
	router.HandleFunc("/", c.IndexGet).Methods("GET")

	// login routes
	router.HandleFunc("/login", c.LoginGet).Methods("GET")
	router.HandleFunc("/login", c.LoginPost).Methods("POST")

	// registration routes
	router.HandleFunc("/registration", c.RegisterGet).Methods("GET")
	router.HandleFunc("/registration", c.RegisterPost).Methods("POST")

	// asset routes
	router.HandleFunc("/assets/css/{file}", c.CSSGet).Methods("GET")
	router.HandleFunc("/assets/js/{file}", c.JSGet).Methods("GET")
}
