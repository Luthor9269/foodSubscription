package databaseConnection

import "time"

// Preference Table struct
type PlaylistTable struct {
	Id                  int       `json:"id"`
	Name                int       `json:"name"`
	SpecialInstructions string    `json:"specialInstructions"`
	DeliveryTime        time.Time `json:"deliveryTime"`
	DeliveryDate        []int     `json:"deliveryDays"` //arr of ints that specify the days of the week
	InPlay              bool      `json:"inPlay"`
	Preferenceid        int       `json:"preferenceid"`
	Userid              int       `json:"userid"`
}
