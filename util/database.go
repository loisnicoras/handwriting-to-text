package util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	// Setup the database (create table if not exists)
	if err := setupDatabase(db); err != nil {
		db.Close() // Close the connection before returning the error
		return nil, err
	}

	return db, nil
}
