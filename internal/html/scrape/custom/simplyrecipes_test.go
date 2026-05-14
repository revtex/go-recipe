package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewSimplyRecipesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.SimplyRecipesHost)

	scraper, err := custom.NewSimplyRecipesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:       "Elise Bauer",
		Categories:   []string{"Appetizer", "Side Dish", "Snack", "Condiment", "Dip", "Quick and Easy", "Restaurant Favorite", "Avocado", "Dip", "Guacamole", "Mexican", "Snack", "Super Bowl", "TexMex"},
		Cuisine:      []string{"Mexican", "TexMex"},
		Description:  "The best guacamole keeps it simple: just ripe avocados and a handful of flavorful mix-ins. Serve it as a dip at your next party or spoon it on top of tacos for an easy dinner upgrade.",
		ImageURL:     "https://www.simplyrecipes.com/thmb/J4kA2m6jKMgkQwZhG-RYpjZBeFQ=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Guacamole-LEAD-6-2-64cfcca253c8421dad4e3fad830219f6.jpg",
		Ingredients:  []string{"2 ripe avocados", "1/4 teaspoon salt, plus more to taste", "1 tablespoon fresh lime or lemon juice", "2-4 tablespoons minced red onion or thinly sliced green onion", "1-2 serrano (or jalapeño) chiles, stems and seeds removed, minced", "2 tablespoons cilantro (leaves and tender stems), finely chopped", "Pinch freshly ground black pepper", "1/2 ripe tomato, chopped (optional)", "Red radish or jicama slices for garnish (optional)", "Tortilla chips , to serve"},
		Instructions: []string{"Cut the avocados in half. Remove the pit. Score the inside of the avocado with a blunt knife and scoop out the flesh with a spoon. (See How to Cut and Peel an Avocado .) Place in a bowl.", "Using a fork, roughly mash the avocado. Don't overdo it! The guacamole should be a little chunky.", "Sprinkle with salt and lime (or lemon) juice. The acid in the lime juice will provide some balance to the richness of the avocado and will help delay the avocados from turning brown. Add the chopped onion, cilantro, black pepper, and chilis. Chili peppers vary individually in their spiciness. So, start with a half of one chili pepper and add more to the guacamole to your desired degree of heat. Remember that much of this is done to taste because of the variability in the fresh ingredients. Start with this recipe and adjust to your taste.", "If making a few hours ahead, place plastic wrap on the surface of the guacamole and press down to cover it to prevent air reaching it. (The oxygen in the air causes oxidation which will turn the guacamole brown.) Garnish with slices of red radish or jicama strips. Serve with your choice of store-bought tortilla chips or make your own homemade tortilla chips . Refrigerate leftover guacamole up to 3 days. Note: Chilling tomatoes dulls their flavor. So, if you want to add chopped tomato to your guacamole, add just before serving. Did you love the recipe? Let us know with a rating and review!"},
		Language:     "",
		Name:         "How to Make the Best Guacamole",
		Nutrition:    recipe.Nutrition{Calories: 252, CarbohydrateGrams: 16, CholesterolMilligrams: 0, FatGrams: 22, FiberGrams: 11, ProteinGrams: 3, SaturatedFatGrams: 3, ServingSize: "Serves 2-4", SodiumMilligrams: 144, SugarGrams: 2, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:     []string{"Condiment", "Dip", "Quick and Easy", "Restaurant Favorite", "Avocado", "Guacamole", "Mexican", "Snack", "Super Bowl", "TexMex", "Gluten-Free", "Low Carb", "Paleo", "Vegan", "Vegetarian", "Appetizer", "Side Dish", "Favorite Summer", "Game Day"},
		PrepTime:     10 * time.Minute,
		Ratings:      4.9,
		RatingsCount: 40,
		SiteName:     "Simply Recipes",
		TotalTime:    10 * time.Minute,
		Yields:       "4",
	}

	scraperTest.Run(t, scraper)
}
