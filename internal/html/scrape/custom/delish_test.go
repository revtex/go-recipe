package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewDelishScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.DelishHost)

	scraper, err := custom.NewDelishScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:       "Lena Abraham",
		Categories:   []string{"low-carb", "low-cost", "30-minute meals", "weeknight meals", "dinner", "main dish"},
		Cuisine:      []string{"Asian Cuisine", "East Asian Cuisine", "Chinese Cuisine", "Vegetarian Diet Cuisine"},
		Description:  "These low-carb egg roll bowls are not only keto and paleo, but are a quick and easy dinner recipe the whole family will love.",
		ImageURL:     "https://hips.hearstapps.com/hmg-prod/images/egg-roll-bowls-index-651c37089e664.jpg?crop=0.500xw:1.00xh;0.256xw,0&resize=1200:*",
		Ingredients:  []string{"1 tbsp. vegetable oil", "1 garlic clove, finely chopped", "1 tbsp. finely chopped peeled ginger", "1 lb. ground pork", "1 tbsp. toasted sesame oil", "1/4 head of green cabbage, thinly sliced", "1/2 yellow onion, thinly sliced", "1 c. shredded carrot", "1/4 c. reduced-sodium soy sauce", "1 tbsp. sriracha", "Kosher salt", "1 scallion, thinly sliced", "1 tbsp. toasted sesame seeds"},
		Instructions: []string{"In a large skillet over medium heat, heat vegetable oil. Add garlic and ginger and cook, stirring, until fragrant, about 1 minute.", "Add pork and cook, breaking up meat into small pieces with a wooden spoon, until browned in parts and cooked through, 8 to 10 minutes.", "Push pork to the side and pour in sesame oil. Add cabbage, onion, and carrot then stir to combine with meat. Add soy sauce and sriracha. Cook, stirring frequently, until cabbage is tender, 5 to 8 minutes; season with salt.", "Divide pork mixture among bowls. Top with scallions and sesame seeds."},
		Language:     "",
		Name:         "Egg Roll Bowls",
		Nutrition:    recipe.Nutrition{Calories: 419, CarbohydrateGrams: 7, CholesterolMilligrams: 82, FatGrams: 32, FiberGrams: 3, ProteinGrams: 22, SaturatedFatGrams: 10, ServingSize: "", SodiumMilligrams: 766, SugarGrams: 4, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:     []string{"content-type: Expanded Recipe", "locale: US", "displayType: recipe", "shortTitle: Egg Roll Bowls", "contentId: 4ba7f257-27c1-45f6-bcd2-6bedead33182", "subsection: Recipes", "collection: Editors' Favorites", "collection: z.Health POC DO NOT USE!!!!!!", "collection: Lunch", "collection: Most Popular", "collection: Delish 100 Best: Dinners", "collection: Budget Recipes", "collection: Dinners", "isSyndicated: false", "eggroll bowls, homemade eggrolls, eggroll filling, healthy asian food, homemade chinese food, easy dinner recipes, healthy bowl recipes,", "NUTRITION: low-carb", "NUTRITION: low-cost", "OCCASION: 30-minute meals", "OCCASION: weeknight meals", "CATEGORY: dinner", "CATEGORY: main dish", "TOTALTIME: 00:35:00", "FILTERTIME: <1HR"},
		Ratings:      4.6148148,
		RatingsCount: 135,
		SiteName:     "Delish",
		PrepTime:     10 * time.Minute,
		CookTime:     25 * time.Minute,
		TotalTime:    35 * time.Minute,
		Yields:       "4 serving(s)",
	}

	scraperTest.Run(t, scraper)
}
