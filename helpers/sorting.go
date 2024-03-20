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

func SortByWeightedScore(restaurants []Restaurant) {
	maxDistance, maxRating, maxPriceStars := getMaxEverything(restaurants)

	sort.Slice(restaurants, func(i, j int) bool {
		leftScore := GetWeightedScore(restaurants[i], maxDistance, maxRating, float64(maxPriceStars))
		rightScore := GetWeightedScore(restaurants[j], maxDistance, maxRating, float64(maxPriceStars))
		return leftScore < rightScore
	})
}

func GetWeightedScore(restaurant Restaurant, maxDistance, maxRating, maxPriceStars float64) float64 {
	distanceWeight := 0.4
	ratingWeight := 0.3
	priceStarsWeight := 0.3

	var weightedScore float64
	weightedScore += distanceWeight * (1 - restaurant.Distance/maxDistance)
	weightedScore += ratingWeight * (restaurant.Rating / maxRating)
	weightedScore += priceStarsWeight * (float64(restaurant.PricingStars) / maxPriceStars)

	return weightedScore
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
