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
		"Pizza", "Burger", "Sushi", "Pasta",
		"Sandwich", "Taco", "Steak", "Soup", "Fried Chicken",
		"Ramen", "Curry", "Tiramisu",
		"Hot Dog", "Shrimp Scampi", "Chicken Parmesan", "Pad Thai", "Caesar Salad",
		"Pho", "Lasagna", "Grilled Cheese", "Fish and Chips",
		"Mashed Potatoes", "Fajitas", "Beef Stroganoff", "Chicken Tikka Masala", "Cobb Salad", "Fruit salad",
	}
	imageURLs := map[string]string{
		"Pizza":                "https://www.simplyrecipes.com/thmb/KE6iMblr3R2Db6oE8HdyVsFSj2A=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2019__09__easy-pepperoni-pizza-lead-3-1024x682-583b275444104ef189d693a64df625da.jpg",
		"Burger":               "https://www.washingtonpost.com/wp-apps/imrs.php?src=https://arc-anglerfish-washpost-prod-washpost.s3.amazonaws.com/public/M6HASPARCZHYNN4XTUYT7H6PTE.jpg&w=1440",
		"Sushi":                "https://www.pbs.org/food/files/2012/09/Sushi-1-1.jpg",
		"Pasta":                "https://food.fnr.sndimg.com/content/dam/images/food/fullset/2021/02/05/Baked-Feta-Pasta-4_s4x3.jpg.rend.hgtvcom.1280.1280.suffix/1615916524567.jpeg",
		"Sandwich":             "https://www.southernliving.com/thmb/UW4kKKL-_M3WgP7pkL6Pb6lwcgM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Ham_Sandwich_011-1-49227336bc074513aaf8fdbde440eafe.jpg",
		"Taco":                 "https://food.fnr.sndimg.com/content/dam/images/food/fullset/2012/2/29/0/0149359_Making-Taco_s4x3.jpg.rend.hgtvcom.1280.1280.suffix/1371603491866.jpeg",
		"Steak":                "https://iamafoodblog.b-cdn.net/wp-content/uploads/2021/02/how-to-cook-steak-1061w.jpg",
		"Soup":                 "https://www.seriouseats.com/thmb/DvSDZoMw8WSOQFAMgf3L2wlfY9Y=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/053123_TomatoSoup-MPPSoupsAndStews-Morgan-Hunt-Glaze-f59a081d7efb4625a75a1a907a6b1cbf.jpg",
		"Fried Chicken":        "https://www.foodandwine.com/thmb/JMrJBrYh3fxDRgkV24_8dZH_zpQ=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/FAW-recipes-crispy-buttermilk-fried-chicken-hero-04-3a32f9a4a1984ecab79fb28e93d4bc00.jpg",
		"Ramen":                "https://www.justonecookbook.com/wp-content/uploads/2023/04/Spicy-Shoyu-Ramen-8055-I.jpg",
		"Curry":                "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/113777-0b21d44.jpg?quality=90&resize=400,363",
		"Tiramisu":             "https://bakeplaysmile.com/wp-content/uploads/2022/06/tiramisu-square.jpg",
		"Hot Dog":              "https://www.seriouseats.com/thmb/QJZXQHDXBfTiUSKstQ1uskJc31g=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/SEA-best-grilled-hot-dogs-recipe-hero-02-9d245c0d43874a3da13a7228682b0dce.jpg",
		"Shrimp Scampi":        "https://assets.bonappetit.com/photos/58a4e12a9fda6d7fbc740e91/1:1/w_3333,h_3333,c_limit/shrimp-scampi.jpg",
		"Chicken Parmesan":     "https://thecozycook.com/wp-content/uploads/2022/08/Chicken-Parmesan-Recipe-f-500x500.jpg",
		"Pad Thai":             "https://tastesbetterfromscratch.com/wp-content/uploads/2018/08/Pad-Thai-Web-7-500x500.jpg",
		"Caesar Salad":         "https://assets.bonappetit.com/photos/624215f8a76f02a99b29518f/1:1/w_2800,h_2800,c_limit/0328-ceasar-salad-lede.jpg",
		"Pho":                  "https://omnivorescookbook.com/wp-content/uploads/2016/01/1511_Easy-Vietnamese-Pho-Noodle-Soup_002-1.jpg",
		"Lasagna":              "https://images.ctfassets.net/hhv516v5f7sj/1AHBnS81eRgSBQwnkub3kF/f6a5d243c45d122ba4d419456c931d5c/lasagnabeeflite_1000x1000.jpg?fm=webp&q=90&w=3840",
		"Grilled Cheese":       "https://cdn.loveandlemons.com/wp-content/uploads/2023/01/grilled-cheese.jpg",
		"Fish and Chips":       "https://www.thespruceeats.com/thmb/sdVTq0h7xZvJjPr6bE2fhh5M3NI=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/SES-best-fish-and-chips-recipe-434856-hero-01-27d8b57008414972822b866609d0af9b.jpg",
		"Mashed Potatoes":      "https://www.skinnytaste.com/wp-content/uploads/2017/11/Garlic-Mashed-Potatoes-7.jpg",
		"Fajitas":              "https://chocolatewithgrace.com/wp-content/uploads/2024/01/CWG-Chicken-Fajitas-5-1-of-1-scaled.jpg",
		"Beef Stroganoff":      "https://www.gimmesomeoven.com/wp-content/uploads/2020/10/Beef-Stroganoff-Recipe-9.jpg",
		"Chicken Tikka Masala": "https://www.seriouseats.com/thmb/DbQHUK2yNCALBnZE-H1M2AKLkok=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/chicken-tikka-masala-for-the-grill-recipe-hero-2_1-cb493f49e30140efbffec162d5f2d1d7.JPG",
		"Cobb Salad":           "https://www.herwholesomekitchen.com/wp-content/uploads/2021/02/cobbsaladrecipe-1.jpg",
		"Fruit salad":          "https://www.spendwithpennies.com/wp-content/uploads/2023/06/1200-Easy-Fresh-Fruit-Salad-SpendWithPennies.jpg",
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
