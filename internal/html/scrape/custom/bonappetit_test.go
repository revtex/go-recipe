package custom_test

import (
	"testing"

	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewBonAppetitScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.BonAppetitHost)

	scraper, err := custom.NewBonAppetitScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:       "Sarah Jampel",
		Description:  "Turn store-bought ravioli into the most sophisticated thing you'll cook this week.",
		ImageURL:     "https://assets.bonappetit.com/photos/62422c4809618646959d68ce/16:9/w_4191,h_2357,c_limit/0328-ravioli-creamy-pea-sauce-lede.jpg",
		Ingredients:  []string{"Kosher salt", "1 10-oz. package frozen peas (about 2 cups)", "1 cup (lightly packed) basil leaves, plus more torn for serving", "½ cup finely grated Parmesan, plus more for serving", "4 Tbsp. unsalted butter, divided", "16–20 oz. fresh or frozen ravioli", "¼ cup coarsely chopped raw pistachios or walnuts", "1 tsp. Aleppo-style pepper", "½ lemon"},
		Instructions: []string{"Place one 10-oz. package frozen peas (about 2 cups) in a fine-mesh sieve and place in a large pot of boiling salted water; cook until peas are tender, about 4 minutes. Add 1 cup (lightly packed) basil leaves and cook until just wilted, about 10 seconds. Lift sieve from water to drain peas and basil and transfer to a blender. (Alternatively, you can skip the sieve and use a spider or slotted spoon to fish out the peas and basil.) Reduce heat to medium-low and keep cooking liquid warm.", "Add ½ cup finely grated Parmesan, 2 Tbsp. unsalted butter, cut into 4 pieces, and ½ cup cooking liquid to blender and blend, gradually increasing speed to high and adding up to ¼ cup additional cooking liquid as needed, until you have a mostly smooth, fairly loose sauce; season with kosher salt.", "Return cooking liquid to a boil over medium-high heat. Add 16–20 oz. fresh (or frozen) ravioli and cook, stirring gently to unstick, until tender, about 5 minutes or according to package directions. Drain carefully; reserve pot.", "Meanwhile, melt remaining 2 Tbsp. unsalted butter in a medium skillet over medium-low heat. Add ¼ cup coarsely chopped raw pistachios or walnuts and cook, stirring often, until butter begins to smell toasty and turn brown, about 5 minutes. Transfer to a small bowl. Add 1 tsp. Aleppo-style pepper, finely grate in zest of ½ lemon, and season lightly with salt; mix well. Slice lemon into wedges.", "Return cooked ravioli to pot, pour pea sauce over, and stir gently to coat. Using a large spoon, transfer ravioli to plates. Top with more Parmesan and torn basil, then spoon buttered nuts over. Serve with lemon wedges for squeezing over."},
		Keywords:     []string{"dinner", "weeknight", "vegetarian", "healthyish", "easy", "quick", "simmer", "ravioli", "pasta & noodles", "pasta", "basil", "pistachio", "walnut", "web"},
		Language:     "",
		Name:         "Weeknight-Fancy Ravioli With Creamy Peas",
		Ratings:      4.6,
		RatingsCount: 34,
		SiteName:     "Bon Appétit",
		Yields:       "2–4 servings",
	}

	scraperTest.Run(t, scraper)
}
