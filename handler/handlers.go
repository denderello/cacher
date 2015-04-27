package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
	m  *sync.Mutex
	db db.KeyValueDatabase
}

func NewSetKeyHandler(m *sync.Mutex, db db.KeyValueDatabase) *SetKeyHandler {
	return &SetKeyHandler{
		m:  m,
		db: db,
	}
}

func (h *SetKeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]

	h.m.Lock()

	if err := h.db.Set(k, v); err != nil {
		http.Error(w, "Could not store key", 500)
		log.Printf("Error while storing key '%s' with value '%s': %#v", k, v, err)
		return
	}

	h.m.Unlock()

	fmt.Fprintf(w, "OK")
}
