# Go API Server with Static File Serving

This repository contains a working example of a Go REST API server that serves static files, created in response to a Reddit post where a user was having trouble serving static HTML files.

## Background

This project was created to help debug an issue posted by [aphroditelady13V on Reddit](https://www.reddit.com/r/golang/comments/1mfmfvk/handler_wont_service_file/). The user was experiencing problems where their API server was finding and attempting to serve static HTML files, but the browser was showing a 404 error instead of the expected content.

After recreating their code structure and implementing mock handlers for the user and product endpoints, I was able to successfully serve the static web page, suggesting I still don't know what is going on.

## Project Structure

```
api_test/
├── api-server.go       # Main API server implementation
├── main.go            # Application entry point
├── example.db         # SQLite database file
├── static/            # Static web assets
│   ├── main.html      # Login page HTML
│   └── favicon.ico    # Favicon
├── user/              # Mock User API handlers
│   └── user.go
└── product/           # Mock Product API handlers
    └── product.go
```

## Features

- **REST API Endpoints**: 
  - `GET /api/v1/users` - Get all users
  - `POST /api/v1/users` - Create a user
  - `GET /api/v1/products` - Get all products  
  - `POST /api/v1/products` - Create a product

- **Static File Serving**:
  - Root path (`/`) serves the main HTML login page
  - `/static/` path serves static assets (CSS, JS, images, etc.)
  - `/favicon.ico` serves the favicon


1. **Access the application**:
This project I served it up on port 3030 (the original user mentioned 8080 was already in use). You can access the application at:
   - Main page: http://localhost:3030
   - API endpoints: http://localhost:3030/api/v1/users or http://localhost:3030/api/v1/products
   - Static files: http://localhost:3030/static/

### Static File Serving Setup

The server uses multiple approaches to serve static content:

```go
// Serve static files from /static/ path
router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

// Serve main HTML page at root
router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/main.html")
})

// Handle favicon
router.Handle("/favicon.ico", http.FileServer(http.Dir("./static")))
```

### The Issue

The original Reddit post described a situation where:
- The server logs showed it was finding and attempting to serve the file
- The absolute path resolution was working correctly
- But the browser displayed a 404 error instead of the HTML content

## License

This project is provided as-is for educational and debugging purposes. Main api-server code written by aphroditelady13V on Reddit.
