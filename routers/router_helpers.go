package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	databaseConnection "github.com/Luthor9269/foodSubscription.git/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurantItemsHandler(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		restaurantItems := getRestaurantItemsHandler(db)
		if restaurantItems == nil || len(restaurantItems) == 0 {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			numItems := len(restaurantItems)
			if numItems > 10 {
				numItems = 10
			}
			first10Items := restaurantItems[:numItems]
			ctx.IndentedJSON(http.StatusOK, first10Items)
		}
	}
}

func getRestaurantItemsHandler(db *sql.DB) []databaseConnection.Item {
	rows, err := db.Query("SELECT items FROM restaurants;")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	var restaurantItems []databaseConnection.Item
	for rows.Next() {
		// Create a variable to hold the JSON or other data structure representing items
		var itemsJSON string
		// Scan the row into the itemsJSON variable
		err := rows.Scan(&itemsJSON)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}

		// Unmarshal the JSON or parse the data structure to extract the items
		var items []databaseConnection.Item
		err = json.Unmarshal([]byte(itemsJSON), &items)
		if err != nil {
			fmt.Println("Error unmarshalling items:", err)
			return nil
		}
		restaurantItems = append(restaurantItems, items...)
	}

	return restaurantItems
}

func PostPreferences(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var preferences databaseConnection.PreferencesTable

		if err := c.ShouldBindJSON(&preferences); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("INSERT INTO preferences (userid, cuisine, minBudget, maxBudget, minRestaurantRating, specialInstructions, restrictions, tags) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			preferences.UserId, preferences.Cuisine, preferences.MinBudget, preferences.MaxBudget, preferences.MinRestaurantRating, preferences.SpecialInstructions, preferences.Restrictions, preferences.Tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with success status
		c.JSON(http.StatusCreated, gin.H{"message": "Preferences added successfully"})
	}
}
