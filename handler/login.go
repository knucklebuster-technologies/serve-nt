package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/qawarrior/secrets"

	"github.com/qawarrior/serve-nt/model"
)

type login struct {
	users *model.UsersCollection
}

func (h *login) get(w http.ResponseWriter, r *http.Request) {
	p := model.PageData{
		Timestamp: time.Now(),
		AppName:   cfg.AppName,
	}
	serveTemplate(w, "./assets/templates/login.html", p)
}

func (h *login) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := &model.User{}
	err := fDecoder.Decode(u, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	txtPwd := u.Password
	u, err = h.users.FindOne(map[string]interface{}{"email": u.Email})
	if err != nil {
		cfg.Logger.Error.Println("User does not Exist")
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}

	if secrets.ComparePassword(txtPwd, u.Password) == false {
		err = errors.New("Email and Password dont match")
		cfg.Logger.Error.Println("Email and Password dont match")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionStore.Get(r, "SNT-SESSION")
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.IsNew {
		session.Values["authenticated"] = true
	}
	err = session.Save(r, w)
	if err != nil {
		cfg.Logger.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rurl := `/profile/` + u.ID.Hex()
	http.Redirect(w, r, rurl, http.StatusSeeOther)
}
