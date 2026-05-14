package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewHelloFreshScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.HelloFreshHost)

	scraper, err := custom.NewHelloFreshScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:      "Sara Heilman",
		Categories:  []string{"main course"},
		Cuisine:     []string{"North American"},
		Description: "Charred broccoli drizzled with lemony Caesar dressing, sprinkled with crunchy garlic panko, and finished with generous shavings of Parmesan cheese is the star of this sheet pan show! It's served with a juicy roasted chicken cutlet and baked sweet potato burnished in the oven to sweet tenderness and dolloped with sour cream.",
		ImageURL:    "https://img.hellofresh.com/f_auto,fl_lossy,h_640,q_auto,w_1200/hellofresh_s3/image/67fcf4921025888df6f416a3-4231d943-7038c46c.jpeg",
		Ingredients: []string{
			"3 tablespoon Sour Cream",
			"1 tablespoon Fry Seasoning",
			"1 unit Seafood Stock Concentrate",
			"3 ounce Caesar Dressing",
			"1 unit Lemon",
			"1 unit Sweet Potato",
			"1 clove Garlic",
			"1 unit Broccoli",
			"1 unit Parmesan Cheese Block",
			"12 ounce Chicken Cutlets",
			"¼ cup Panko Breadcrumbs",
			"1 teaspoon Garlic Powder",
			"teaspoon (tsp) Salt",
			"4 teaspoon (tsp) Olive Oil",
			"1 teaspoon (tsp) Cooking Oil",
			"1 tablespoon (tbsp) Butter",
			"teaspoon (tsp) Black Pepper",
		},
		Instructions: []string{
			"Adjust racks to top and middle positions and preheat oven to 425 degrees. Wash and dry produce.",
			"Cut broccoli into four wedges. (TIP: For even pieces that hold together, place the broccoli floret side down and cut through the stem. It's OK if some of the florets fall off; set them aside to roast too.) Halve sweet potato lengthwise. Pierce cut sides of sweet potato with a fork 4-5 times. Peel and mince or grate garlic. Zest and quarter lemon.",
			"Place broccoli on one side of a baking sheet. Toss with a large drizzle of olive oil, garlic powder, salt, and pepper.",
			"On empty side of same sheet, coat sweet potato with a large drizzle of olive oil; season all over with salt and pepper. Arrange cut sides down.",
			"Roast on top rack until sweet potato is tender and broccoli is tender and lightly charred, 25-30 minutes.",
			"Pat chicken* dry with paper towels and drizzle with oil. Season all over with Fry Seasoning, salt, and pepper.",
			"Place chicken on a second baking sheet; roast on middle rack until golden brown and cooked through, 20-25 minutes.",
			"Meanwhile, in a small bowl, whisk together Caesar dressing, 1 tsp stock concentrate, lemon zest, and juice from one lemon wedge (2 tsp stock concentrate and juice from two lemon wedges for 4 servings). TIP: If you prefer a milder Caesar dressing, start with ¼ tsp stock concentrate or simply omit it!",
			"Melt 1 TBSP butter (2 TBSP for 4 servings) in a small pan over medium-high heat. Add garlic and cook, stirring, until fragrant, 30-60 seconds.",
			"Add panko; season with salt and pepper. Cook, stirring occasionally, until panko is golden brown and toasted, 3-5 minutes. Turn off heat and set aside.",
			"Thinly slice chicken crosswise.",
			"Divide chicken, sweet potato, and broccoli between plates. Top sweet potato with as much sour cream as you like. Drizzle broccoli with as much Caesar dressing as you like; top with garlic panko. Using a vegetable peeler, shave Parmesan over broccoli. Serve with any remaining lemon wedges and any remaining dressing on the side.",
		},
		Name: "Roasted Chicken & Crispy Caesar Broccoli with Sweet Potato, Sour Cream & Parmesan",
		Nutrition: recipe.Nutrition{
			Calories:              820,
			CarbohydrateGrams:     46,
			CholesterolMilligrams: 175,
			FatGrams:              49,
			FiberGrams:            7,
			ProteinGrams:          49,
			SaturatedFatGrams:     14,
			SodiumMilligrams:      1110,
			SugarGrams:            11,
		},
		Ratings:      5.0,
		RatingsCount: 3,
		SiteName:     "HelloFresh",
		TotalTime:    45 * time.Minute,
	}

	scraperTest.Run(t, scraper)
}
