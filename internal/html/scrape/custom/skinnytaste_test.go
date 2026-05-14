package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewSkinnyTasteScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.SkinnyTasteHost)

	scraper, err := custom.NewSkinnyTasteScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Categories:   []string{"Dinner"},
		CookTime:     360 * time.Minute,
		Cuisine:      []string{"Mexican"},
		Description:  "Crock Pot Chicken Taco Chili is an easy slow cooker dump recipe using freezer and pantry staples! This has been one of my most popular recipes since 2008!",
		Instructions: []string{"Combine beans, onion, chili peppers, corn, tomato sauce, diced tomato, cumin, chili powder and taco seasoning in a slow cooker and mix well.", "Nestle the chicken in to completely cover and cook on LOW for 8 to 10 hours or on HIGH for 4 to 6 hours.", "Half hour before serving, remove chicken and shred.", "Return chicken to slow cooker and stir in.", "Top with fresh cilantro and your favorite toppings!"},
		ImageURL:     "https://www.skinnytaste.com/wp-content/uploads/2008/11/crockpot-chicken-taco-chili-8.jpg",
		Ingredients:  []string{"1 small onion (chopped)", "15.5 oz can black beans, drained", "15.5 oz can kidney beans, drained", "8 ounce can tomato sauce", "10 oz package frozen corn kernels", "2 cans diced tomatoes w/chilies (10 ounces each)", "4 oz can chopped green chili peppers (chopped)", "1 packet reduced sodium taco seasoning or homemade (see below)", "1 tbsp cumin", "1 tbsp chili powder", "24 oz 3 boneless skinless chicken breasts", "1/4 cup chopped fresh cilantro", "1 1/2 tablespoons cumin", "1 1/2 tablespoons chili powder", "1/4 teaspoon garlic powder", "1/4 teaspoon onion powder", "1/4 teaspoon dried oregano", "1/2 teaspoon paprika", "1 teaspoon kosher salt", "1/2 teaspoon black pepper"},
		Language:     "",
		Name:         "Crock Pot Chicken Taco Chili",
		Nutrition:    recipe.Nutrition{Calories: 220, CarbohydrateGrams: 28, CholesterolMilligrams: 44, FatGrams: 3, FiberGrams: 8.5, ProteinGrams: 21, SaturatedFatGrams: 0, ServingSize: "1 scant cup", SodiumMilligrams: 729, SugarGrams: 6, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:     []string{"chili recipe", "crock pot chicken recipes", "crock pot chicken taco chili", "crock pot chili", "how to make slow cooker chicken taco chili", "slow cooker chicken taco chili", "slow cooker dump chicken taco chili"},
		PrepTime:     5 * time.Minute,
		Ratings:      4.91,
		RatingsCount: 240,
		SiteName:     "Skinnytaste",
		TotalTime:    360 * time.Minute,
		Yields:       "10 servings",
	}

	scraperTest.Run(t, scraper)
}
