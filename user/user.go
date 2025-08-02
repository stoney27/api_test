package user

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Store struct {
	db *sql.DB
}

type Handler struct {
	store *Store
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.getUsers).Methods("GET")
	router.HandleFunc("/users", h.createUser).Methods("POST")
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"users": []}`))
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created"}`))
}
