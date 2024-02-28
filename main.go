package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	handler "github.com/loisnicoras/handwriting-to-text/handlers"
)

func connectToDB(username, password, hostname, dbname string) (*sql.DB, error) {
	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)

	// Open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database connection: %w", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		db.Close() // Close the connection before returning the error
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func main() {
	apiKey := flag.String("apiKey", "", "Specify an API key")
	addr := flag.String("addr", "8080", "Specify the server address and port")
	dbUser := flag.String("dbUser", "lois", "Specify the database user")
	dbPass := flag.String("dbPass", "emanuel", "Specify the database password")
	dbHost := flag.String("dbHost", "localhost", "Specify the database hostname")
	dbName := flag.String("dbName", "my_database", "Specify the database name")
	flag.Parse()

	db, err := connectToDB(*dbUser, *dbPass, *dbHost, *dbName)
	if err != nil {
		fmt.Printf("failed to connect to the db %s", err)
		return
	}

	r := chi.NewRouter()

	r.Get("/", handler.HomeHandler)
	r.Post("/signup", handler.CreateUser(db))
	r.Post("/extract-text", handler.UploadHandler(apiKey))

	r.Route("/exercises", func(r chi.Router) {
		r.Get("/", handler.GetExercises(db))
		r.Get("/{exerciseID}", handler.GetExercise(db))
		r.Post("/{exerciseID}", handler.SubmitExercise(db))
	})

	fmt.Printf("Server is running at http://localhost:%s\n", *addr)
	if err := http.ListenAndServe(":"+*addr, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
