package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewSeriousEatsScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.SeriousEatsHost)

	scraper, err := custom.NewSeriousEatsScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:      "Sho Spaeth",
		Categories:  []string{"Entree", "Side Dish", "Appetizers and Hors d'Oeuvres", "Quick and Easy", "Sides", "Snacks"},
		CookTime:    40 * time.Minute,
		Cuisine:     []string{"Asian", "Korean"},
		Description: "Crispy-chewy rice, gooey cheese, funky and sour kimchi...what's not to like?",
		ImageURL:    "https://www.seriouseats.com/thmb/FpZlQan7XTk0pEn-PMLDCay598w=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__serious_eats__seriouseats.com__2019__10__20191014-cheesy-rice-vicky-wasik-13-b353b816b4f147cca230eed39cc193fe.jpg",
		Ingredients: []string{
			"1 tablespoon (15g) unsalted butter, softened",
			"4 cups cooked short-grain rice (see note)",
			"3 tablespoons (45ml) gochujang",
			"3 tablespoons (45ml) soy sauce",
			"1 tablespoon (15ml) rice vinegar",
			"4 thinly sliced scallions (~80g), white and green parts divided",
			"3 ounces (85g) grated low-moisture mozzarella cheese (see note), divided",
			"2 ounces (57g) grated Gruyère cheese (see note), divided",
			"5 ounces (141g) chopped drained kimchi (see note)",
			"1 ounce (28g) grated Cotija cheese (see note)",
		},
		Instructions: []string{
			"Preheat oven to 400°F (200°C). Grease a 10-inch cast iron pan with 1 tablespoon (15g) butter, making sure to fully cover both the bottom and the sides.",
			"In a medium mixing bowl, combine rice, gochujang, soy sauce, rice vinegar, and sliced scallion whites. Using a flexible spatula, mix thoroughly.",
			"Scrape half of rice mixture into buttered cast iron skillet and, using the bottom of a drinking glass or measuring cup, press down firmly to create a single even layer of seasoned rice. Distribute half of the mozzarella and Gruyère over the layer of rice, then scrape the rest of the rice mixture over grated cheese. Using the bottom of a drinking glass or measuring cup, press down firmly to create an even top layer of seasoned rice. Transfer pan to oven and cook for 35 minutes.",
			"Remove pan from oven. Turn off oven and turn broiler on high. While broiler preheats, top rice with chopped kimchi. Distribute remaining Gruyère and mozzarella over kimchi and sprinkle Cotija over the other cheeses.",
			"Place pan under broiler for about 2 minutes, or until Gruyère and mozzarella melt and bubble and Cotija begins to char in spots. Remove pan from under broiler, top with sliced scallion greens, and serve immediately.",
		},
		Keywords: []string{"kimchi"},
		Name:     "Crispy Kimchi Cheese Rice Recipe",
		Nutrition: recipe.Nutrition{
			Calories:              544,
			CarbohydrateGrams:     73,
			CholesterolMilligrams: 57,
			FatGrams:              19,
			FiberGrams:            1,
			ProteinGrams:          19,
			SaturatedFatGrams:     11,
			ServingSize:           "Serves 4 as a side dish or snack, or 1 as a late-night feast",
			SodiumMilligrams:      1892,
			SugarGrams:            8,
		},
		PrepTime:     15 * time.Minute,
		Ratings:      4.9,
		RatingsCount: 7,
		SiteName:     "Serious Eats",
		TotalTime:    55 * time.Minute,
		Yields:       "4",
	}

	scraperTest.Run(t, scraper)
}
