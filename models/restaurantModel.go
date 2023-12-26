package databaseConnection

import "time"

// Preference Table struct
type RestaurantTable struct {
	Id             int       `json:"id"`
	Cuisines       []int     `json:"cuisines"`
	Address        string    `json:"address"`
	OpeningTime    time.Time `json:"openingTime"`
	ClosingTime    time.Time `json:"closingTime"`
	RestaurantName string    `json:"restaurantName"`
}
