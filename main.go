package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	handler "github.com/loisnicoras/handwriting-to-text/handlers"
)

func main() {
	apiKey := flag.String("apiKey", "", "Specify an api key")
	flag.Parse()
	// Create a new Chi router
	router := chi.NewRouter()

	// Define handlers for different routes
	router.Get("/", handler.HomeHandler)
	router.Post("/upload", handler.UploadHandler(apiKey))

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
