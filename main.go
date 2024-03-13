package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	handler "github.com/loisnicoras/handwriting-to-text/handlers"
	login "github.com/loisnicoras/handwriting-to-text/login"

)

func connectToDB(username, password, hostname, port, dbname string) (*sql.DB, error) {
	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

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
	dbPort := flag.String("dbPort", "3306", "Specify the database port")
	dbName := flag.String("dbName", "my_database", "Specify the database name")
	projectId := flag.String("projectId", "moonlit-shadow-325207", "Specify the google cloud project id")
	region := flag.String("region", "us-central1", "Specify the vertex ai region")

	flag.Parse()

	db, err := connectToDB(*dbUser, *dbPass, *dbHost, *dbPort, *dbName)
	if err != nil {
		fmt.Printf("failed to connect to the db %s", err)
		return
	}

	r := chi.NewRouter()

	r.Get("/", handler.HomeHandler)
	r.Get("/login", login.HandleGoogleLogin)
	r.Get("/callback", login.HandleGoogleCallback(db))
	r.Post("/extract-text", handler.UploadHandler(apiKey))

	r.Route("/exercises", func(r chi.Router) {
		// r.Use(handler.AuthMiddleware)
		r.Get("/", handler.GetExercises(db))
		r.Get("/{exerciseID}", login.AuthMiddleware(handler.GetExercise(db)))
		r.Post("/{exerciseID}", login.AuthMiddleware(handler.SubmitExercise(db, *projectId, *region)))
	})

	fmt.Printf("Server is running at http://localhost:%s\n", *addr)
	if err := http.ListenAndServe(":"+*addr, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
