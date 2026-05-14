package recipe

import (
	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"

	"github.com/PuerkitoBio/goquery"
)

type recipeScraperFunc func(*goquery.Document) (recipe.Scraper, error)

var hostToScraper = map[string]recipeScraperFunc{
	custom.AllRecipesHost:         custom.NewAllRecipesScraper,
	custom.BBCGoodFoodHost:        custom.NewBBCGoodFoodScraper,
	custom.BonAppetitHost:         custom.NewBonAppetitScraper,
	custom.BudgetBytesHost:        custom.NewBudgetBytesScraper,
	custom.CookieAndKateHost:      custom.NewCookieAndKateScraper,
	custom.DelishHost:             custom.NewDelishScraper,
	custom.EpicuriousHost:         custom.NewEpicuriousScraper,
	custom.Food52Host:             custom.NewFood52Scraper,
	custom.FoodNetworkHost:        custom.NewFoodNetworkScraper,
	custom.ForksOverKnivesHost:    custom.NewForksOverKnivesScraper,
	custom.HalfBakedHarvestHost:   custom.NewHalfBakedHarvestScraper,
	custom.HelloFreshHost:         custom.NewHelloFreshScraper,
	custom.LoveAndOtherSpicesHost: custom.NewLoveAndOtherSpicesScraper,
	custom.MarthaStewartHost:      custom.NewMarthaStewartScraper,
	custom.MinimalistBakerHost:    custom.NewMinimalistBakerScraper,
	custom.NYTimesHost:            custom.NewNYTimesScraper,
	custom.PinchOfYumHost:         custom.NewPinchOfYumScraper,
	custom.SeriousEatsHost:        custom.NewSeriousEatsScraper,
	custom.SimplyRecipesHost:      custom.NewSimplyRecipesScraper,
	custom.SkinnyTasteHost:        custom.NewSkinnyTasteScraper,
	custom.TasteOfHomeHost:        custom.NewTasteOfHomeScraper,
	custom.TastyHost:              custom.NewTastyScraper,
	custom.TheKitchnHost:          custom.NewTheKitchnScraper,
	custom.TheSpruceEatsHost:      custom.NewTheSpruceEatsScraper,
}
