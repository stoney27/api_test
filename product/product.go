package product

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
	router.HandleFunc("/products", h.getProducts).Methods("GET")
	router.HandleFunc("/products", h.createProduct).Methods("POST")
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"products": []}`))
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Product created"}`))
}
