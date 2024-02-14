package main

import (
	"fmt"
	"net/http"
	handler "github.com/loisnicoras/handwriting-to-text/handlers"
	"github.com/go-chi/chi"
)

func main() {
	// Create a new Chi router
	router := chi.NewRouter()

	// Define handlers for different routes
	router.Get("/", handler.HomeHandler)
	router.Post("/upload", handler.UploadHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
