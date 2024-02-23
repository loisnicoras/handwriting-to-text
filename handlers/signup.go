package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Database credentials
const (
	username = "lois" // Assuming you are using root user
	password = "emanuel"
	hostname = "localhost" // Docker service name
	dbPort   = "3306"
	dbname   = "my_database"
)

func connectToDB() (*sql.DB, error) {
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
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	_, err = db.Exec("USE " + dbname)
	if err != nil {
		return nil, fmt.Errorf("failed to use this database: %w", err)
	}

	// Create table 'users' in the database
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
            email VARCHAR(100) NOT NULL,
            password VARCHAR(100) NOT NULL
        )
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDB()
	if err != nil {
		http.Error(w, "failed to connect to the db", http.StatusInternalServerError)
		return
	}

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Insert data into MySQL table
	_, err = db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)", first_name, last_name, email, password)
	if err != nil {
		http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
		return
	}

	// Return a success message
	fmt.Fprintf(w, "Data inserted successfully!")
}
