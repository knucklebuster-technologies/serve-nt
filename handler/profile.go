package handler

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/qawarrior/serve-nt/model"
)

type profile struct {
	users  *model.UsersCollection
	events *model.EventsCollection
}

func (h *profile) get(w http.ResponseWriter, r *http.Request) {
	if !authenicated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// obtain and validate the id
	id := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(id) == false {
		cfg.Logger.Error.Println("The ID:", id, "is not valid")
		http.Error(w, "Invalid Id", http.StatusUnauthorized)
		return
	}

	oid := bson.ObjectIdHex(id)
	u, err := h.users.FindOne(map[string]interface{}{"_id": oid})
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	page := model.PageData{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}
	p := model.ProfileData{
		page,
		*u,
	}
	serveTemplate(w, "./assets/templates/profile.html", p)
}
