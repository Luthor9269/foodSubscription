package routers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func SetupRouter(db *sql.DB) {
	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	router.GET("/restaurant-items", GetRestaurantItemsHandler(db))
	router.POST("/preference", PostPreferences(db))
	//

	err = router.Run("localhost:8080")
	fmt.Println("router up")
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
