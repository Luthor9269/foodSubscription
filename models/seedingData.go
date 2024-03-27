package databaseConnection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
)

type Restaurant struct {
	ID             int
	UniqueID       string
	Items          []Item
	Location       string
	OperatingHours string
	Name           string
}

type Item struct {
	Name        string
	Price       float64
	Description string
	ImageURL    string
}

func seedingData(db *sql.DB) error {

	restaurants := []Restaurant{
		{
			UniqueID:       "234567",
			Items:          generateItems(12),
			Location:       "123 Orchard Rd, Singapore",
			OperatingHours: "Mon-Sun: 11am-11pm",
			Name:           "Singaporean Delights",
		},
		{
			UniqueID:       "890123",
			Items:          generateItems(14),
			Location:       "456 Marina Bay Sands, Singapore",
			OperatingHours: "Mon-Sat: 10am-10pm, Sun: 12pm-8pm",
			Name:           "Seafood Paradise",
		},
		{
			UniqueID:       "345678",
			Items:          generateItems(16),
			Location:       "789 Sentosa Cove, Singapore",
			OperatingHours: "Mon-Sun: 12pm-12am",
			Name:           "Hawker Haven",
		},
		{
			UniqueID:       "456789",
			Items:          generateItems(13),
			Location:       "321 Raffles Ave, Singapore",
			OperatingHours: "Mon-Sun: 9am-10pm",
			Name:           "Raffles Bistro",
		},
		{
			UniqueID:       "567890",
			Items:          generateItems(12),
			Location:       "678 Little India, Singapore",
			OperatingHours: "Mon-Sun: 8am-9pm",
			Name:           "Spice Village",
		},
		{
			UniqueID:       "678901",
			Items:          generateItems(15),
			Location:       "987 Clarke Quay, Singapore",
			OperatingHours: "Mon-Sun: 11:30am-12am",
			Name:           "Riverside Grill",
		},
		{
			UniqueID:       "789012",
			Items:          generateItems(14),
			Location:       "876 Chinatown, Singapore",
			OperatingHours: "Mon-Sun: 10am-11pm",
			Name:           "Lucky Dragon",
		},
	}

	for _, restaurant := range restaurants {
		err := SeedRestaurant(db, restaurant)
		if err != nil {
			fmt.Println("Error seeding table")
			log.Fatal("Error seeding restaurant:", err)

		}
	}
	fmt.Println("Seeding completed successfully")
	// returning no error
	return nil
}

func generateItems(numItems int) []Item {
	foodNames := []string{
		"Pizza", "Burger", "Sushi", "Pasta", "Salad",
		"Sandwich", "Taco", "Steak", "Soup", "Fried Chicken",
		"Ramen", "Curry", "Tiramisu", "Ice Cream", "Cheesecake",
		"Hot Dog", "Shrimp Scampi", "Chicken Parmesan", "Pad Thai", "Caesar Salad",
		"Pho", "Lasagna", "Grilled Cheese", "Hamburger", "Fish and Chips",
		"Mashed Potatoes", "Fajitas", "Beef Stroganoff", "Chicken Tikka Masala", "Cobb Salad",
	}
	imageURLs := map[string]string{
		"Pizza":           "../images/pizza.png",
		"Burger":          "../images/hamburger.png",
		"Sushi":           "../images/fruits.png",
		"Ramen":           "../images/provolone.png",
		"Lasagna":         "../images/provolone.png",
		"Mashed Potatoes": "../images/fruits.png",
	}

	descriptions := []string{
		"A classic dish with a perfect blend of flavors and textures",
		"Juicy and flavorful, guaranteed to satisfy your cravings",
		"Melt-in-your-mouth goodness that will leave you wanting more",
		"Savory and aromatic, a true comfort food experience",
		"Bold and spicy, sure to tantalize your taste buds",
		"Fresh and invigorating, perfect for a light and healthy meal",
		"Tender and succulent, a delight for meat lovers",
		"Indulgent and decadent, a treat for your senses",
		"Crispy on the outside, tender on the inside, simply irresistible",
		"Sweet and satisfying, a perfect ending to any meal",
		"Rich and creamy, a taste of luxury in every bite",
		"Homemade goodness, made with love and care",
		"A gourmet delight, crafted with the finest ingredients",
		"Hearty and filling, a meal that sticks to your ribs",
		"Succulent and juicy, cooked to perfection",
		"Fragrant and aromatic, a feast for the senses",
		"Smokey and flavorful, reminiscent of a summer barbecue",
		"Wholesome and nourishing, a healthy choice for any occasion",
		"Exquisite and elegant, fit for a special occasion",
		"Decadent and indulgent, a sinful pleasure",
		"Irresistible and addictive, you won't be able to get enough",
		"Authentic flavors, straight from the heart of [country/region]",
		"Nutritious and delicious, a guilt-free treat",
		"Fresh and vibrant, bursting with flavor",
		"A culinary masterpiece, expertly crafted for your enjoyment",
		"Comfort food at its finest, guaranteed to warm your soul",
		"A symphony of flavors, each bite is a delight",
		"Tender and mouthwatering, a true culinary sensation",
		"Zesty and refreshing, a burst of flavor in every bite",
		"A delicious fusion of [cuisine] and [cuisine], a unique culinary experience",
	}

	//imageFolder := "../images/"

	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		// Generate a random number between 0 and 29
		randomNumber := rand.Intn(30)
		foodName := foodNames[randomNumber]
		//imagePath := imageFolder + foodName + ".png" // Assuming image file names match food names
		items[i] = Item{
			Name:        fmt.Sprintf(foodName),
			Price:       float64(randomNumber) * 2.5,
			Description: fmt.Sprintf(descriptions[randomNumber]),
			ImageURL:    imageURLs[foodName],
		}
	}
	return items
}

func SeedRestaurant(db *sql.DB, restaurant Restaurant) error {

	// Convert items slice to JSON string
	itemsJSON, err := json.Marshal(restaurant.Items)
	if err != nil {
		return err
	}

	var seedRestaurantRow = `
INSERT INTO restaurants 
    (unique_id, location, operating_hours, name, items) 
VALUES ($1, $2, $3, $4, $5) RETURNING id`
	// Execute the SQL statement with placeholders
	_, err = db.Exec(seedRestaurantRow, restaurant.UniqueID, restaurant.Location, restaurant.OperatingHours, restaurant.Name, string(itemsJSON))
	if err != nil {
		return err
	}
	return nil
}
