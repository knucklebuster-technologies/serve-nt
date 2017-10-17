package handler

import (
	"encoding/json"
	"net/http"

	"github.com/qawarrior/serve-nt/model"
)

type events struct {
	collection *model.EventsCollection
}

func (h *events) get(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("servee")
	if q != "" {
		m, err := h.collection.Find(map[string]interface{}{"_id": q})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(m)
		return
	}

	m, err := h.collection.Find(map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(m)
}

func (h *events) post(w http.ResponseWriter, r *http.Request) {
	m := &model.Event{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m, err = h.collection.Insert(m)
	json.NewEncoder(w).Encode(m)
}
