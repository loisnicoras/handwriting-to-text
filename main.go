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
	addr := flag.String("addr", "8080", "Specify the server address and port")
	flag.Parse()

	r := chi.NewRouter()

	r.Get("/", handler.HomeHandler)
	r.Post("/signup", handler.CreateUser)
	r.Post("/upload", handler.UploadHandler(apiKey))

	r.Route("/exercises", func(r chi.Router) {
		r.Get("/", handler.GetExercises)
		r.Get("/{exerciseID}", handler.GetExercise)
		r.Post("/{exerciseID}", handler.SubmitExercise)
	})

	fmt.Printf("Server is running at http://localhost:%s\n", *addr)
	http.ListenAndServe(":"+*addr, r)
}
