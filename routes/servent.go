package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/controllers"
	mgo "gopkg.in/mgo.v2"
)

func servent(dbname string, db *mgo.Session, router *mux.Router) {
	c, err := controllers.NewServent(db.Copy().DB(dbname))
	if err != nil {
		return
	}
	router.HandleFunc("/v1/servent", c.Create).Methods("POST")
	router.HandleFunc("/v1/servent", c.Read).Methods("GET")
	router.HandleFunc("/v1/servent", c.Update).Methods("PUT")
	router.HandleFunc("/vi/servent", c.Delete).Methods("DELETE")
}
