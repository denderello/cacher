package handler

import (
	"fmt"
	"net/http"

	"github.com/denderello/cacher/db"
	"github.com/gorilla/mux"
)

type GetKeyHandler struct {
	db db.KeyValueDatabase
}

func NewGetKeyHandler(db db.KeyValueDatabase) *GetKeyHandler {
	return &GetKeyHandler{
		db: db,
	}
}

func (h *GetKeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	v, err := h.db.Get(vars["key"])
	if err != nil {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, v)
}

type SetKeyHandler struct {
	db db.KeyValueDatabase
}

func NewSetKeyHandler(db db.KeyValueDatabase) *SetKeyHandler {
	return &SetKeyHandler{
		db: db,
	}
}

func (h *SetKeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]

	err := h.db.Set(k, v)
	if err != nil {
		http.Error(w, "Could not store key", 500)
	}
	fmt.Fprintf(w, "OK")
}
