package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/controllers"
	mgo "gopkg.in/mgo.v2"
)

func event(dbname string, db *mgo.Session, router *mux.Router) {
	c, err := controllers.NewEvent(db.Copy().DB(dbname))
	if err != nil {
		return
	}
	router.HandleFunc("/v1/event", c.Create).Methods("POST")
	router.HandleFunc("/v1/event", c.Read).Methods("GET")
	router.HandleFunc("/v1/event", c.Update).Methods("PUT")
	router.HandleFunc("/vi/event", c.Delete).Methods("DELETE")
}
