package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const FoodNetworkHost = "foodnetwork.com"

// NewFoodNetworkScraper returns a new instance of FoodNetworkScraper.
func NewFoodNetworkScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &FoodNetworkScraper{schema: s}, nil
}

// FoodNetworkScraper is a custom recipe scraper for foodnetwork.com.
type FoodNetworkScraper struct {
	schema *schema.RecipeScraper
}

func (m *FoodNetworkScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *FoodNetworkScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *FoodNetworkScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *FoodNetworkScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *FoodNetworkScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *FoodNetworkScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *FoodNetworkScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *FoodNetworkScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *FoodNetworkScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *FoodNetworkScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *FoodNetworkScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *FoodNetworkScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *FoodNetworkScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *FoodNetworkScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *FoodNetworkScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *FoodNetworkScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *FoodNetworkScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *FoodNetworkScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *FoodNetworkScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *FoodNetworkScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *FoodNetworkScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *FoodNetworkScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
