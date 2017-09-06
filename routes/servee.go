package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/controllers"
	mgo "gopkg.in/mgo.v2"
)

func servee(dbname string, db *mgo.Session, router *mux.Router) {
	c, err := controllers.NewServee(db.Copy().DB(dbname))
	if err != nil {
		return
	}
	resource := "/servees"
	resourceid := resource + "/{id}"
	router.HandleFunc(resource, c.Create).Methods("POST")
	router.HandleFunc(resource, c.Read).Methods("POST")
	router.HandleFunc(resourceid, c.Read).Methods("GET")
	router.HandleFunc(resourceid, c.Update).Methods("PUT")
	router.HandleFunc(resourceid, c.Delete).Methods("DELETE")
}
