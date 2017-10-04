package handlers

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/models"
)

func profileidGet(w http.ResponseWriter, r *http.Request) {
	// obtain and validate the id
	id := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(id) == false {
		cfg.Logger.Error.Println("The ID:", id, "is not valid")
		http.Error(w, "Invalid Id", http.StatusUnauthorized)
		return
	}

	// create user and find by id
	u := models.NewUser()
	u.ID = bson.ObjectIdHex(id)
	err := u.FindByID()
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, "User not foumd", http.StatusUnauthorized)
		return
	}

	p := profiledata{
		pagedata{
			Timestamp: time.Now(),
			AppName:   cfg.AppName,
		},
		*u,
	}
	serveTemplate(w, "./assets/templates/profile.html", p)
}
