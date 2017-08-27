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
	router.HandleFunc("/v1/servee", c.Create).Methods("POST")
	router.HandleFunc("/v1/servee", c.Read).Methods("GET")
	router.HandleFunc("/v1/servee", c.Update).Methods("PUT")
	router.HandleFunc("/vi/servee", c.Delete).Methods("DELETE")
}
