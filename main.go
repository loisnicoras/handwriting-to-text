package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	handler "github.com/loisnicoras/handwriting-to-text/handlers"
)

func main() {
	apiKey := flag.String("apiKey", "", "Specify an API key")
	addr := flag.String("addr", ":8080", "Specify the server address and port")
	flag.Parse()

	// Create a new Chi router
	router := chi.NewRouter()

	// Define handlers for different routes
	router.Get("/", handler.HomeHandler)
	router.Post("/upload", handler.UploadHandler(apiKey))

	// Start the server
	fmt.Printf("Server is running at http://localhost:%s\n", *addr)
	http.ListenAndServe(":"+*addr, router)
}
