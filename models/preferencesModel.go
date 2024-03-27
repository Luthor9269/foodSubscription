package databaseConnection

// Preference Table struct
type PreferencesTable struct {
	Id                  int     `json:"id"`
	UserId              int     `json:"userid"`
	Cuisine             []int   `json:"cuisine"`
	MinBudget           float32 `json:"minBudget"`
	MaxBudget           float32 `json:"maxBudget"`
	MinRestaurantRating float32 `json:"minRestaurantRating"`
	SpecialInstructions string  `json:"specialInstructions"`
	Restrictions        []int   `json:"restrictions"`
}
