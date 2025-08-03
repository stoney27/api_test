package main

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stoney27/api_test/product"
	"github.com/stoney27/api_test/user"
)

//go:embed static/*
var staticFiles embed.FS

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	// Create a sub-filesystem for static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal("Error creating static filesystem:", err)
	}

	// Serve static files from embedded filesystem
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve main.html from embedded filesystem
		data, err := staticFiles.ReadFile("static/main.html")
		if err != nil {
			log.Println("Error reading main.html:", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("NOT FOUND:", r.URL.Path)
		http.NotFound(w, r)
	})

	// Serve favicon from embedded filesystem
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		data, err := staticFiles.ReadFile("static/favicon.ico")
		if err != nil {
			log.Println("Error reading favicon.ico:", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
