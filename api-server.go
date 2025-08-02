package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stoney27/api_test/product"
	"github.com/stoney27/api_test/user"
)

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

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}
	log.Println("Current working directory:", cwd)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		absPath, err := filepath.Abs("./static/main.html")
		if err != nil {
			log.Println("Error getting absolute path:", err)
		} else {
			log.Println("Serving file from:", absPath)
		}
		http.ServeFile(w, r, "./static/main.html")

		// log.Println("Serving / with static HTML")
		//   w.Header().Set("Content-Type", "text/html")
		//   w.WriteHeader(http.StatusOK)
		//   w.Write([]byte(`<html><body><h1>Hello from Go server!</h1></body></html>`))
	})

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("NOT FOUND:", r.URL.Path)
		http.NotFound(w, r)
	})

	router.Handle("/favicon.ico", http.FileServer(http.Dir("./static")))

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
