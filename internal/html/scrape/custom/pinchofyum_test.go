package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewPinchOfYumScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.PinchOfYumHost)

	scraper, err := custom.NewPinchOfYumScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        "Lindsay Ostrom",
		Categories:    []string{"Sandwich"},
		CookTime:      10 * time.Minute,
		CookingMethod: "Layer",
		Cuisine:       []string{"American"},
		Description:   "Pillowy sun-dried tomato focaccia gets piled with pesto mayo, turkey, juicy tomato slices, and greens. This sandwich is a beauty and practically a lifestyle!",
		ImageURL:      "https://pinchofyum.com/tachyon/Trader-Joes-Sun-Dried-Tomato-Focaccia-Turkey-Sandwich-Square.png?fit=225%2C225",
		Ingredients: []string{
			"1 loaf of Trader Joe’s sun-dried tomato focaccia",
			"one 7-ounce package deli turkey",
			"1 large tomato, thinly sliced",
			"one handful of greens – sprouts, microgreens, spinach, etc.",
			"a few pieces of thinly sliced red onion",
			"1/4 cup mayo",
			"1/4 cup pesto (I like the Costco brand pesto)",
		},
		Instructions: []string{
			"Preheat the oven to 400 degrees. Toast the focaccia in the oven for 5-8 minutes.",
			"Mix the pesto and the mayo together in a small bowl to make a pesto mayo.",
			"Keeping the bread flat on a cutting board, turn your serrated knife sideways and slice horizontally through the center of the loaf. (You don’t want to flip the top and lose those yummy crumbles.) Keeping the top piece flat, set it aside.",
			"Layer the sandwich with pesto mayo, turkey, tomato, greens, red onion, and finish with dollops of pesto mayo. I dollop the top layer of pesto mayo directly on top of the greens and onions so that I don’t have to flip my top piece upside down to spread it on the bread.",
			"Gently press your top piece back on top of the sandwich. When it’s nice and secure, use a large knife to cut through the sandwich. I usually cut it down the middle and across 3 times for a total of 6 pieces. Serve and enjoy! This is sandwich goals.",
		},
		Keywords: []string{"turkey sandwich", "trader joe's sandwich recipe", "focaccia sandwich"},
		Name:     "Trader Joe's Sun-Dried Tomato Focaccia Turkey Sandwich",
		Nutrition: recipe.Nutrition{
			Calories:              338,
			CarbohydrateGrams:     40.1,
			CholesterolMilligrams: 18.2,
			FatGrams:              12.8,
			FiberGrams:            2.7,
			ProteinGrams:          15.8,
			SaturatedFatGrams:     2.2,
			SodiumMilligrams:      840,
			SugarGrams:            4.6,
		},
		PrepTime:     10 * time.Minute,
		Ratings:      5,
		RatingsCount: 46,
		SiteName:     "Pinch of Yum",
		TotalTime:    20 * time.Minute,
		Yields:       "6 servings",
	}

	scraperTest.Run(t, scraper)
}
