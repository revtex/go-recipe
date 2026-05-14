package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewHalfBakedHarvestScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.HalfBakedHarvestHost)

	scraper, err := custom.NewHalfBakedHarvestScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Categories:  []string{"Main Course"},
		CookTime:    20 * time.Minute,
		Cuisine:     []string{"American", "Greek"},
		Description: "Grilled chicken, buttery rice pilaf, veggies, and Tzatziki. A fresh, easy Greek-inspired dinner bowl!",
		ImageURL:    "https://www.halfbakedharvest.com/wp-content/uploads/2025/07/Chicken-Souvlaki-Rice-Pilaf-1-scaled.jpg",
		IngredientGroups: []recipe.IngredientGroup{
			{
				Purpose: "",
				Ingredients: []string{
					"▢ 2 pounds chicken breasts, cut into small pieces",
					"▢ 6 tablespoons salted butter, melted",
					"▢ 2 tablespoons lemon juice",
					"▢ 6 cloves garlic, chopped",
					"▢ 1 tablespoon smoked paprika",
					"▢ 2 teaspoons dried oregano",
					"▢ 1 teaspoon dried rosemary (or thyme)",
					"▢ sea salt and pepper",
					"▢ chili flakes",
					"▢ 2-3 cups bell pepper quarters/zucchini rounds",
					"▢ 1 cup cherry tomatoes",
					"▢ 1/2 cup crumbled feta cheese",
					"▢ Tzatziki, for serving",
					"▢ mixed herbs - basil, thyme, dill for serving",
				},
			},
			{
				Purpose: "Rice Pilaf",
				Ingredients: []string{
					"▢ 2 tablespoons salted butter",
					"▢ 1 cup basmati rice",
					"▢ 3/4 cup orzo",
					"▢ 2-3 tablespoons pine nuts (optional)",
				},
			},
		},
		Ingredients: []string{
			"2 pounds chicken breasts, cut into small pieces",
			"6 tablespoons salted butter, melted",
			"2 tablespoons lemon juice",
			"6 cloves garlic, chopped",
			"1 tablespoon smoked paprika",
			"2 teaspoons dried oregano",
			"1 teaspoon dried rosemary ((or thyme))",
			"sea salt and pepper",
			"chili flakes",
			"2-3 cups bell pepper quarters/zucchini rounds",
			"1 cup cherry tomatoes",
			"1/2 cup crumbled feta cheese",
			"Tzatziki, for serving",
			"mixed herbs - basil, thyme, dill for serving",
			"2 tablespoons salted butter",
			"1 cup basmati rice",
			"3/4 cup orzo",
			"2-3 tablespoons pine nuts ((optional))",
		},
		Instructions: []string{
			"1. To make the rice. Melt the butter with the rice, orzo, and pine nuts in a medium pot over medium heat. Let the butter brown. Add 2 cups water or broth (or bone broth). Bring to a boil over high heat. Reduce the heat to low, cover, and simmer for 10 minutes. Turn the heat off. Allow the rice to cook, covered, for 15-20 minutes more. Fluff the rice with a fork.2. To make the chicken. Mix the chicken, butter, lemon juice, garlic, smoked paprika, oregano, rosemary, salt, pepper, and chili flakes in a bowl. If time allows, let it sit for 10 minutes or up to overnight to deepen the flavor.3. Set your grill, grill pan, or skillet to medium-high heat. Thread the chicken onto skewers. Grill the skewers until lightly charred and cooked through, turning them occasionally throughout cooking, about 10 to 12 minutes total. During the same time, grill the peppers and zucchini for 3-4 minutes per side. 4. Serve the chicken with Tzatziki sauce and crumbled feta over bowls of rice pilaf. Arrange the gilled veggies and tomatoes on the side. Top with herbs. I like to add naan or pita bread too!",
		},
		Keywords: []string{"grill", "skewers", "under one hour"},
		Name:     "Quick Chicken Souvlaki Rice Pilaf",
		Nutrition: recipe.Nutrition{
			Calories:    572,
			ServingSize: "1 serving",
		},
		PrepTime:     20 * time.Minute,
		Ratings:      4.42,
		RatingsCount: 12,
		SiteName:     "Half Baked Harvest",
		TotalTime:    40 * time.Minute,
		Yields:       "6",
	}

	scraperTest.Run(t, scraper)
}
