package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewForksOverKnivesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.ForksOverKnivesHost)

	scraper, err := custom.NewForksOverKnivesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Categories:  []string{"Vegan Baked & Stuffed Recipes"},
		CookTime:    55 * time.Minute,
		ImageURL:    "https://www.forksoverknives.com/uploads/Spelt-Banana-Walnut-Bread-wordpress.jpg?format=webp&optimize=high&precrop=16%3A9%2Csmart",
		Ingredients: []string{"½ cup chopped walnuts", "1½ cups mashed ripe banana", "¾ cup pitted dates", "¾ cup unsweetened, unflavored almond milk", "¼ cup unsweetened applesauce", "2 teaspoons pure vanilla extract", "1 cup spelt flour", "1 cup brown rice flour", "2 teaspoons ground cinnamon", "1½ teaspoons regular or sodium-free baking powder", "½ teaspoon baking soda", "½ teaspoon sea salt"},
		Instructions: []string{
			"Preheat oven to 350°F. Line four mini (5½ x 3½-inch) loaf pans with parchment paper. Spread walnuts in a shallow baking pan and bake 10 minutes or until toasted and fragrant. Let cool.",
			"In a blender combine the next five ingredients (through vanilla). Cover and blend until smooth. In a large bowl whisk together the remaining ingredients. Add banana mixture to flour mixture. Stir to combine. Fold in all but 2 tablespoons of the walnuts. Divide batter among prepared loaf pans. Sprinkle with the remaining 2 tablespoons of walnuts.",
			"Bake 40 to 45 minutes or until a toothpick inserted in the center comes out clean. Cool completely in pans on a wire rack.",
		},
		Name: "Spelt Banana Bread with Toasted Walnuts",
		Nutrition: recipe.Nutrition{
			Calories:          259,
			CarbohydrateGrams: 49,
			FatGrams:          5.7,
			FiberGrams:        5.9,
			ProteinGrams:      6,
			SaturatedFatGrams: 0.7,
			ServingSize:       "½ mini loaf",
			SodiumMilligrams:  307,
			SugarGrams:        15,
		},
		PrepTime:     20 * time.Minute,
		Ratings:      5.0,
		RatingsCount: 8,
		SiteName:     "Forks Over Knives",
		TotalTime:    1*time.Hour + 15*time.Minute,
		Yields:       "4 mini loaves",
	}

	scraperTest.Run(t, scraper)
}
