package routers

//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	_ "github.com/go-sql-driver/mysql"
//	"net/http"
//)
//
//func GetRestaurantItemsHandler(c *gin.Context) {
//	username, _ := c.Get("db_username")
//	dbname, _ := c.Get("db_name")
//	restaurantItems := getRestaurantItemsHandler(username, dbname)
//	if restaurantItems == nil || len(restaurantItems) == 0 {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.IndentedJSON(http.StatusOK, restaurantItems)
//	}
//}
//
//func getRestaurantItemsHandler(username, dbname string) []models.RestaurantItem {
//	db, err := sql.Open("mysql", username+"@tcp(127.0.0.1:3306)/"+dbname)
//
//	if err != nil {
//		fmt.Println("Err", err.Error())
//		return nil
//	}
//
//	defer db.Close()
//
//	results, err := db.Query("SELECT * FROM restaurantItems")
//
//	if err != nil {
//		fmt.Println("Err", err.Error())
//		return nil
//
//	}
//	restaurantItems := []RestaurantItem{}
//	for results.Next() {
//		var i Item
//		err = results.Scan(&i.Id, &i.Title, &i.Director)
//
//		if err != nil {
//			panic(err.Error())
//		}
//		restaurantItems = append(restaurantItems, i)
//	}
//	return restaurantItems
//}
