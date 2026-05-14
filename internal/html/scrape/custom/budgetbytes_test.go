package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewBudgetBytesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.BudgetBytesHost)

	scraper, err := custom.NewBudgetBytesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Categories:   []string{"Dinner", "Main Course"},
		CookTime:     20 * time.Minute,
		Cuisine:      []string{"Italian"},
		Description:  "This budget-friendly Chicken Marsala recipe is so rich and flavorful, you won't believe it didn't come from your favorite Italian restaurant!",
		ImageURL:     "https://www.budgetbytes.com/wp-content/uploads/2024/03/Chicken-Marsala-Close.jpg",
		Ingredients:  []string{"2 chicken breasts (boneless, skinless, $2.98)", "½ tsp salt (divided, $0.02)", "1 tsp black pepper (freshly cracked, $0.05)", "1 Tbsp Italian seasoning ($0.30)", "½ cup all-purpose flour (1 Tbsp reserved, $0.04*)", "3 Tbsp vegetable oil (divided, $0.12)", "1 tsp salt ($0.03)", "8 oz. baby bella mushrooms (sliced, (roughly 2 cups) $2.17**)", "½ white onion (sliced thinly, $0.58)", "½ cup Marsala cooking wine ($2.33***)", "1 cup chicken broth ($0.13****)", "2 Tbsp fresh Italian parsley (minced, $0.12)", "2 Tbsp salted butter ($0.27)"},
		Instructions: []string{"Cut the chicken breasts in half lengthwise. Place one chicken breast at a time in a heavy-duty freezer bag (unsealed) and use a tenderizing mallet (or a rolling pin…whatever is safe and sturdy that you have on hand) to hammer each one 1/4\" thin, starting in the thicker middle part of the breasts and working your way to the ends. Repeat this process until all breast pieces are 1/4\" thick and equally flat.", "Coat: Combine ½ tsp salt, black pepper, dried Italian herb seasoning, and all-purpose flour in a shallow bowl, reserving 1 Tbsp of flour. Lightly coat the chicken breasts in the seasoning mixture and shake off excess flour. Transfer to a clean plate.", "Cook: Heat vegetable oil in a large skillet over medium-high heat. Once hot, add the chicken breasts and cook on each side for 3-4 minutes. Try to not crowd your pan! You can do this in 2 batches if you need to. Once the chicken is cooked, transfer to a clean plate and cover with foil to keep warm.", "Sauté: Turn the heat down to medium. Add 1 tsp sea salt, sliced mushrooms, and onions. Spread them out on the pan, flipping as needed so each mushroom piece browns, about 3 minutes. Scrape the bottom of your pan with a wooden spoon as you go; the browned bits on the bottom become a tasty part of your sauce!", "Add the Marsala wine and broth to the pan, scraping the browned bits off the bottom again. Simmer until the liquid reduces by 1/3, about 4 minutes on medium-high heat.", "Thicken: Now, make a slurry with your reserved 1 Tbsp of flour; whisk the flour with a small amount of water or liquid from your cooking pan until there are no lumps. Then, whisk in the flour \u201cslurry\u201d, minced parsley, and butter until it has completely melted. Finally, add the chicken cutlets back, flipping them halfway, cooking for another 4 minutes total. Let the sauce thicken while the chicken finishes cooking in the sauce.", "Serve: Garnish with additional minced parsley (optional) and adjust salt and pepper to your liking. Enjoy over any pasta you have on hand!"},
		Language:     "",
		Name:         "Chicken Marsala",
		Nutrition:    recipe.Nutrition{Calories: 407, CarbohydrateGrams: 21, CholesterolMilligrams: 0, FatGrams: 19, FiberGrams: 2, ProteinGrams: 29, SaturatedFatGrams: 0, ServingSize: "1 serving", SodiumMilligrams: 1280, SugarGrams: 0, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		Keywords:     []string{"chicken marsala"},
		PrepTime:     15 * time.Minute,
		Ratings:      4.84,
		RatingsCount: 12,
		SiteName:     "Budget Bytes",
		TotalTime:    35 * time.Minute,
		Yields:       "4 servings",
	}

	scraperTest.Run(t, scraper)
}
