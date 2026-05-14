package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewMarthaStewartScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.MarthaStewartHost)

	scraper, err := custom.NewMarthaStewartScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:      "Martha Stewart",
		Categories:  []string{"Dinner", "Entree"},
		CookTime:    6*time.Hour + 20*time.Minute,
		Cuisine:     []string{"Mexican", "Tex-Mex"},
		Description: "Our carne guisada recipe uses the slow-cooker, making it easy to get this rich beef stew recipe ready for dinner. Serve it over rice or with tortillas.",
		ImageURL:    "https://www.marthastewart.com/thmb/xaXETOHjCvC7H-_5udF6HzYFti0=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/MS-329032-slow-cooker-carne-guisada-hero-3x2-727cffb3ac1748c497443fae95e8ed9d.jpg",
		Ingredients: []string{
			"2.5 pounds beef chuck roast or bottom round, cut into 1-inch pieces",
			"Coarse salt and ground pepper",
			"2 tablespoons vegetable oil",
			"1 medium white onion, diced medium",
			"1 medium green bell pepper, seeded and diced medium",
			"1 large jalapeno, seeded and diced small",
			"5 garlic cloves, roughly chopped",
			"1.5 teaspoons ground cumin",
			"0.75 teaspoon chili powder",
			"0.75 teaspoon dried oregano",
			"6 tablespoons all-purpose flour",
			"1.75 cups low-sodium chicken broth",
			"1 can (14 ounces) diced tomatoes",
			"2 bay leaves",
			"Flour tortillas, warmed, grated cheddar, and cilantro, for serving",
		},
		Instructions: []string{
			"Brown beef: Season beef with salt and pepper. In a large skillet, heat 2 teaspoons oil over high. In 2 batches, cook beef until browned on all sides, 5 minutes per batch (add 2 teaspoons more oil for second batch). Transfer to a 5- to 6-quart slow cooker.",
			"Cook onion, garlic, and peppers; add seasonings: In same skillet, cook 2 teaspoons oil, onion, bell pepper, jalapeño, and garlic over medium, stirring and scraping up browned bits with a wooden spoon, until vegetables are tender, 5 minutes. Add cumin, chili powder, oregano, and flour and cook 1 minute. Slowly pour broth into skillet, stirring until liquid is smooth. Simmer 2 minutes.",
			"Transfer to slow cooker: Transfer mixture to slow cooker, along with tomatoes and bay leaves. Season to taste with salt and pepper and stir to combine. Cover and cook on high 6 hours. Serve over rice or in tortillas with cheese and cilantro.",
		},
		Keywords:  []string{"Beef chuck roast", "Bottom round steak", "Buffet", "Dinner", "Dinner party", "Main course", "Mexican", "Potluck", "Slow Cook", "Stews", "Super Bowl party", "Tex-Mex", "Tomato sauces"},
		Name:      "Slow-Cooker Carne Guisada",
		PrepTime:  15 * time.Minute,
		SiteName:  "Martha Stewart",
		TotalTime: 6*time.Hour + 35*time.Minute,
		Yields:    "10",
	}

	scraperTest.Run(t, scraper)
}
