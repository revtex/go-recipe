package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewAllRecipesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.AllRecipesHost)

	scraper, err := custom.NewAllRecipesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        "Kathryn Hendrix, RDN, LD",
		Categories:    []string{"Dinner"},
		CookTime:      40 * time.Minute,
		Cuisine:       []string{"American"},
		Description:   "This super easy dump and bake meatball casserole made with frozen meatballs and prepared marinara sauce comes together in minutes and then just needs to bake in the oven. Perfect family staple on weeknights.",
		ImageURL:      "https://www.allrecipes.com/thmb/yIq7SgrJS2Q98jFI5RW9xHmQjnM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/11695447-Dump-and-Bake-Meatball-Casserole-ddmfs-beauty-4x3-17079-311b92f41da848bbbca8d40fb97b8fb4.jpg",
		Ingredients:   []string{"12 ounces dry rotini pasta (3 1/2 cups)", "2 1/2 cups water", "1 (24 ounce) jar marinara sauce", "1 (14 ounce) bag frozen meatballs, thawed", "1 cup chopped bell pepper (1 medium)", "1 teaspoon salt", "2 cups shredded mozzarella cheese (8 ounces), divided", "1/4 cup grated Parmesan cheese", "Chopped fresh herbs, such as parsley or basil, optional"},
		Instructions:  []string{"Gather all ingredients. Preheat the oven to 425 degrees F (220 degrees C).", "Combine pasta, water, marinara sauce, meatballs, peppers, and salt in a 9x13-inch baking dish.", "Cover tightly with aluminum foil and bake until pasta is al dente, 30 to 35 minutes.", "Remove the foil and stir in 1 cup mozzarella cheese. Sprinkle with remaining 1 cup of mozzarella cheese and Parmesan cheese.", "Bake, uncovered, until cheese is melted, 10 minutes more.", "Serve garnished with fresh herbs, if desired."},
		Language:      "",
		Name:          "Dump and Bake Meatball Casserole",
		Nutrition:     recipe.Nutrition{Calories: 362, CarbohydrateGrams: 28, CholesterolMilligrams: 56, FatGrams: 19, FiberGrams: 4, ProteinGrams: 19, SaturatedFatGrams: 8, ServingSize: "", SodiumMilligrams: 1281, SugarGrams: 8, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:      []string{"publisher-tested"},
		PrepTime:      5 * time.Minute,
		Ratings:       4.2,
		RatingsCount:  6,
		SiteName:      "Allrecipes",
		SuitableDiets: nil,
		TotalTime:     45 * time.Minute,
		Yields:        "8",
	}

	scraperTest.Run(t, scraper)
}
