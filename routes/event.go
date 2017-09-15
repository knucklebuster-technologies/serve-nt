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
	resource := "/events"
	router.HandleFunc(resource, c.Create).Methods("POST")
	router.HandleFunc(resource, c.Read).Methods("GET")
	router.HandleFunc(resource+"/title/{title}", c.ReadTitle).Methods("GET")
	router.HandleFunc(resource, c.Update).Methods("PUT")
	router.HandleFunc(resource, c.Delete).Methods("DELETE")
}