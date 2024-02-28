package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		err := r.ParseForm()
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
}
