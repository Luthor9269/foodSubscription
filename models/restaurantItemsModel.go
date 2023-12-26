package databaseConnection

// Preference Table struct
type RestaurantItemsTable struct {
	Id              int     `json:"id"`
	ItemDescription string  `json:"itemDescription"`
	ItemName        string  `json:"itemName"`
	Price           float32 `json:"price"`
	Tags            []int   `json:"tags"`
	RestaurantId    string  `json:"restaurantid"`
}
