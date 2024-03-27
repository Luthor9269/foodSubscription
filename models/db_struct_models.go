package databaseConnection

type RestaurantItem struct {
	ItemId         int     `json:"id"`
	RestaurantName string  `json:"name"`
	RestaurantId   int     `json:"unique_id"`
	ItemName       string  `json:"items"`
	ItemPrice      float64 `json:"price"`
}

type Preference struct {
	PreferenceID        int
	Cuisine             []int
	MinBudget           float64
	MaxBudget           float64
	MinRestaurantRating float64
	SpecialInstruction  *string
	Restrictions        []int
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
