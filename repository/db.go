// Repository package contains repository operations
// for account and transfer models
// This file contains the configurations for the database
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 123
	dbname   = "desafiotecnico"
)

// Start DB
// Initialize the database
// If the operation is successful, returns the connected database
// If the operation fails, returns nil
func StartDB() *sql.DB {

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, user, port, dbname) // DOCKER

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error opening database connection!", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error checking if database connection is still active!", err)
		return nil
	}

	fmt.Println("Database connected!")

	return db
}
