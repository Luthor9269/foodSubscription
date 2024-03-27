package databaseConnection

import (
	"database/sql"
	"fmt"
	"log"
)

// initTables initializes the required database tables 'users' and 'attendance' if they do not exist.
// It takes a database connection as a parameter and executes SQL commands to create the tables.
func initTables(db *sql.DB) error {
	// Queries to execute

	//Creating database if not exists

	var createUserTable = `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY, 
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		phoneNumber VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL,
		sessionid UUID 
	);
	`

	var createRestaurantTable = `
	CREATE TABLE IF NOT EXISTS restaurants(
    id SERIAL PRIMARY KEY,
    unique_id VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    operating_hours VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    items JSON NOT NULL
);
	`

	var createPreferenceTable = `
	CREATE TABLE IF NOT EXISTS preferences(
		id SERIAL PRIMARY KEY,
		userid INT,
		cuisine INT[],
		minBudget FLOAT,
		maxBudget FLOAT,
		minRestaurantRating FLOAT,
		specialInstructions VARCHAR(255),
		restrictions INT[]
	)
	`

	var createPlaylistTable = `
	CREATE TABLE IF NOT EXISTS playlist(
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		specialInstructions VARCHAR(255),
		deliveryTime TIMESTAMP NOT NULL,
		deliverydays INT[] NOT NULL,
		inPlay BOOLEAN NOT NULL,
		preferenceid INT REFERENCES preferences(id),
		userid INT REFERENCES users(id) NOT NULL
	);
	`

	// db.Exec()
	//db.ExecContext--> Can be used to setTimeout,Cancel signal or dealine. So it doesn't run indefinetely
	// Since I did not define a schema first. psql will use a default public schema
	_, err := db.Exec(createUserTable)
	if err != nil {
		fmt.Println("Error adding user table")
		log.Fatalln(err)
	}

	_, err = db.Exec(createRestaurantTable)
	if err != nil {
		fmt.Println("Error adding restaurant table")
		log.Fatalln(err)
	}
	_, err = db.Exec(createPreferenceTable)
	if err != nil {
		fmt.Println("Error adding preference table")
		log.Fatalln(err)
	}
	_, err = db.Exec(createPlaylistTable)
	if err != nil {
		fmt.Println("Error adding playlist table")
		log.Fatalln(err)
	}
	return nil
}
