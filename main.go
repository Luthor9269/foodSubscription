package main

import (
	databaseConnection "github.com/Luthor9269/foodSubscription.git/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	db := databaseConnection.ConnectToDB()
	defer db.Close()
	http.ListenAndServe(":8080", nil)
	//start server
	// http.ListenAndServe(":8080", nil)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
