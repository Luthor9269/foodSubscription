package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

func main() {
	// Creating and setting up a simple server
}

func faker() {
	gofakeit.Seed(0)
	// Generate a fake dish
	dish := struct {
		Name        string  `fake:"{food} {word} Dish"`
		Description string  `fake:"{sentence}"`
		Price       float64 `fake:"{price}"`
	}{}
	gofakeit.Struct(&dish)
	fmt.Printf("Dish: %+v\n", dish)
}
