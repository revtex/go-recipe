package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewBBCGoodFoodScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.BBCGoodFoodHost)

	scraper, err := custom.NewBBCGoodFoodScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:       "Esther Clark",
		Categories:   []string{"Dinner", "Pasta", "Supper"},
		CookTime:     20 * time.Minute,
		Description:  "Cook up this classic sauce in one pan, then toss with spaghetti for a simple midweek meal. It's budget-friendly too, making it a great meal for the family",
		ImageURL:     "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/puttanesca-cfb4e42.jpg?resize=440,400",
		Ingredients:  []string{"3 tbsp olive oil", "1 onion finely chopped", "2 large garlic cloves crushed", "½ tsp chilli flakes (optional)", "400g can chopped tomatoes", "5 anchovy fillets, finely chopped", "120g pitted black olives", "2 tbsp capers drained", "300g dried spaghetti", "½ small bunch of parsley finely chopped"},
		Instructions: []string{"Heat the oil in a non-stick pan over a medium-low heat. Add the onion along with a generous pinch of salt and fry for 10 mins, or until soft. Add the garlic and chilli, if using, and cook for a further minute.", "Stir the tomatoes, anchovies, olives and capers into the onion, bring to a gentle simmer and cook, uncovered, for 15 mins. Season to taste.", "Meanwhile, bring a large pan of salted water to the boil. Cook the spaghetti following pack instructions, then drain and toss with the sauce and parsley."},
		Language:     "",
		Name:         "Spaghetti puttanesca",
		Nutrition:    recipe.Nutrition{Calories: 495, CarbohydrateGrams: 66, CholesterolMilligrams: 0, FatGrams: 19, FiberGrams: 6, ProteinGrams: 13, SaturatedFatGrams: 3, ServingSize: "", SodiumMilligrams: 1.8, SugarGrams: 8, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:     []string{"Anchovies", "Budget", "Capers", "Esther Clark", "Low calorie", "Olives", "Puttanesca", "Spaghetti"},
		PrepTime:     15 * time.Minute,
		Ratings:      5.0,
		RatingsCount: 249,
		SiteName:     "Good Food",
		TotalTime:    35 * time.Minute,
	}

	scraperTest.Run(t, scraper)
}
