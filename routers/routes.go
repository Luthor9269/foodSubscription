package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func SetupRouter(res http.ResponseWriter, req *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		// Set DB credentials in the Gin context
		c.Set("db_username", username)
		c.Set("db_password", password)
		c.Set("db_name", dbname)
		c.Next()
	})

	router.GET("/restaurant-item/{itemId}", GetRestaurantItemsHandler)
	router.GET("/playlist", temporaryHandler)
	router.DELETE("/restaurant-item/{itemId}")
	router.POST("/preference")

	err = router.Run("localhost:8080")
	if err != nil {
		return
	}
}

type Playlist struct {
	PlaylistID     int
	PreferenceID   int
	PlaylistName   string
	DeliveryTiming string
	DeliveryDays   []int
	ItemIDs        []string
	IsPlaying      bool
}

func temporaryHandler(c *gin.Context) {
	playlist := Playlist{
		PlaylistID:     1,
		PreferenceID:   1,
		PlaylistName:   "My Playlist",
		DeliveryTiming: "12:00 PM",
		DeliveryDays:   []int{1, 3, 5},
		ItemIDs:        []string{"item1", "item2", "item3"},
		IsPlaying:      true,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(playlist)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c.IndentedJSON(http.StatusOK, jsonData)
}
