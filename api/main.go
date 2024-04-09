package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	exercise "github.com/loisnicoras/handwriting-to-text/handlers/exercise"
	home "github.com/loisnicoras/handwriting-to-text/handlers/home"
	login "github.com/loisnicoras/handwriting-to-text/handlers/login"
	upload "github.com/loisnicoras/handwriting-to-text/handlers/upload"
	"github.com/rs/cors"
)

func connectToDB(username, password, hostname, port, dbname string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func main() {
	addr := flag.String("addr", "8080", "Specify the server address and port")
	apiKey := flag.String("apiKey", "", "Specify an API key")
	projectId := flag.String("projectId", "moonlit-shadow-325207", "Specify the google cloud project id")
	region := flag.String("region", "us-central1", "Specify the vertex ai region")
	dbUser := flag.String("dbUser", "lois", "Specify the database user")
	dbPass := flag.String("dbPass", "emanuel", "Specify the database password")
	dbHost := flag.String("dbHost", "localhost", "Specify the database hostname")
	dbPort := flag.String("dbPort", "3306", "Specify the database port")
	dbName := flag.String("dbName", "my_database", "Specify the database name")

	flag.Parse()

	db, err := connectToDB(*dbUser, *dbPass, *dbHost, *dbPort, *dbName)
	if err != nil {
		fmt.Printf("failed to connect to the db %s", err)
		return
	}

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	})
	r.Use(cors.Handler)

	r.Get("/", home.HomeHandler)
	r.Get("/login", login.HandleGoogleLogin)
	r.Get("/logout", login.LogOut(db))
	r.Get("/callback", login.HandleGoogleCallback(db))
	r.Get("/user-data", login.GetUserData(db))
	r.Post("/extract-text", upload.UploadHandler(apiKey))

	r.Route("/exercises", func(r chi.Router) {
		r.Get("/", exercise.GetExercises(db))
		r.Get("/{exerciseID}", exercise.GetExercise(db))
		r.Post("/{exerciseID}", login.AuthMiddleware(exercise.SubmitExercise(db, *projectId, *region)))
	})

	fmt.Printf("Server is running at http://localhost:%s\n", *addr)
	if err := http.ListenAndServe(":"+*addr, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
