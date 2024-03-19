package helpers

import "sort"

type Restaurant struct {
	ID            uint64
	Cuisines      []int
	Rating        float64
	Distance      float64
	PricingStars  int
	WeightedScore float64
}

func GetWeightedScores(restaurants []Restaurant) []Restaurant {
	cuisineWeight := 0.4
	distanceWeight := 0.3
	ratingWeight := 0.2
	priceStarsWeight := 0.1

	maxDistance, maxRating, maxPriceStars := getMaxEverything(restaurants)
	// normalize to a value between 0-1
	// higher values are better
	for i := range restaurants {
		restaurants[i].WeightedScore += distanceWeight * (1 - restaurants[i].Distance/maxDistance)
		restaurants[i].WeightedScore += ratingWeight * (restaurants[i].Rating / maxRating)
		restaurants[i].WeightedScore += priceStarsWeight * (float64(restaurants[i].PricingStars) / float64(maxPriceStars))
	}

	sort.Slice(restaurants, func(i, j int) bool {
		return restaurants[i].WeightedScore > restaurants[j].WeightedScore
	})
	return restaurants
}

func getMaxEverything(restaurants []Restaurant) (float64, float64, int) {
	maxDistance := restaurants[0].Distance
	maxRating := restaurants[0].Rating
	maxPriceStars := restaurants[0].PricingStars

	for _, m := range restaurants {
		if m.Distance > maxDistance {
			maxDistance = m.Distance
		}
		if m.Rating > maxRating {
			maxRating = m.Rating
		}
		if m.PricingStars > maxPriceStars {
			maxPriceStars = m.PricingStars
		}
	}
	return maxDistance, maxRating, maxPriceStars
}
