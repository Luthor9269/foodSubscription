package main

import (
	databaseConnection "github.com/Luthor9269/foodSubscription.git/models"
)

func main() {
	db := databaseConnection.ConnectToDB()
	defer db.Close()

	//start server
	// http.ListenAndServe(":8080", nil)
}
