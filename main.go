package main

import (
	"database/sql"
	"log"
)

func main() {
	log.Println("Starting API server...")
	// Initialize database connection (example)
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()
	// Test database connection
	if err := db.Ping(); err != nil {
		log.Println("Warning: Database ping failed:", err)
	} else {
		log.Println("Database connection successful")
	}
	// need to run on 3030 port already had 8080 being used
	apiServer := NewAPIServer(":3030", db)
	log.Printf("About to start server on %s", apiServer.addr)
	if err := apiServer.Run(); err != nil {
		log.Fatal("Error starting API server:", err)
	}
}
