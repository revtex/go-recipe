package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewNYTimesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.NYTimesHost)

	scraper, err := custom.NewNYTimesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:      "Yewande Komolafe",
		Categories:  []string{"dinner", "easy", "for two", "quick", "weeknight", "beans", "one pot", "main course"},
		CookTime:    30 * time.Minute,
		Description: "Like a warm and gentle nudge, masala spice gives onions and chickpeas a distinctively comforting heartiness. Glimmering with droplets of ghee, they become rich in this any-season dish. Tearing the tofu allows for ample crooks and crannies that cradle and accentuate the aromatic goodness of the spice. Cherry tomatoes, slightly and delicately blistered, are welcome as juicy bursts of acidity in every bite. Serve this over rice, or with a gently poached egg, along with a few slices of lime for squeezing.",
		ImageURL:    "https://static01.nyt.com/images/2025/06/13/multimedia/yk-sheet-pan-masala-tofu-vjmb/yk-sheet-pan-masala-tofu-vjmb-videoSixteenByNineJumbo1600.jpg",
		Ingredients: []string{
			"1 (14-ounce) block firm or extra-firm tofu, drained",
			"3 1/2 tablespoons ghee or vegetable oil",
			"Salt and black pepper",
			"1 large red onion, finely chopped, some saved for garnish",
			"1 (2-inch) piece fresh ginger, finely grated",
			"2 garlic cloves, finely grated",
			"1 teaspoon ground tandoori or garam masala (homemade or store-bought)",
			"1 (15-ounce) can chickpeas, rinsed",
			"1 pint cherry tomatoes",
			"1/4 cup fresh mint or dill leaves, chopped, plus more leaves for garnish",
			"Lime wedges, for serving",
		},
		Instructions: []string{
			"Slice the tofu in half horizontally and place on a clean kitchen or paper towel to dry.",
			"Set a 10-inch skillet over medium heat and add 1 tablespoon ghee. Once the ghee begins to shimmer, season both sides of the tofu with salt and pepper, place in the pan and sear without moving until the tofu is browned, about 4 minutes. Turn the pieces over and brown the other side, 4 to 5 minutes more. Transfer the tofu to a plate.",
			"Add 2 tablespoons ghee to the same skillet and heat over medium until shimmering. Add the onion (saving some for garnish) and cook, stirring often, for 4 minutes, or until translucent. Add the ginger, garlic and masala spice and season generously with salt and pepper. Stir for 1 to 2 minutes, or until fragrant.",
			"Stir in the chickpeas and cook for 3 minutes, or until the chickpeas begin to sizzle.",
			"Turn the heat up to medium-high, add the remaining ½ tablespoon ghee, then add the tomatoes and 1 teaspoon salt. Cook without stirring until the tomatoes are just beginning to pop open and the chickpeas are warmed through, 4 to 5 minutes. Stir in the mint.",
			"Break the tofu into 1-inch pieces and toss in the skillet to coat with chickpea-tomato mix. Cook for another 2 to 3 minutes, or until warmed through. Remove from heat, taste and adjust seasonings, if necessary. Garnish with remaining chopped raw onions and a few leaves of fresh mint. Serve with lime wedges for squeezing.",
		},
		Keywords: []string{"budget", "cherry tomato", "chickpea", "extra-firm tofu", "firm tofu", "red onion", "stovetop", "dairy-free", "gluten-free", "healthy", "vegan", "vegetarian"},
		Name:     "Masala Chickpeas With Tofu and Blistered Tomatoes",
		Nutrition: recipe.Nutrition{
			CarbohydrateGrams:   37,
			FatGrams:            20,
			FiberGrams:          10,
			ProteinGrams:        18,
			SaturatedFatGrams:   2,
			SodiumMilligrams:    831,
			SugarGrams:          9,
			UnsaturatedFatGrams: 15,
		},
		PrepTime:     5 * time.Minute,
		Ratings:      5,
		RatingsCount: 2078,
		SiteName:     "NYT Cooking",
		TotalTime:    35 * time.Minute,
		Yields:       "2 to 4 servings",
	}

	scraperTest.Run(t, scraper)
}
