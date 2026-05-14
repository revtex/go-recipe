package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const BBCGoodFoodHost = "bbcgoodfood.com"

// NewBBCGoodFoodScraper returns a new instance of BBCGoodFoodScraper.
func NewBBCGoodFoodScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &BBCGoodFoodScraper{schema: s}, nil
}

// BBCGoodFoodScraper is a custom recipe scraper for bbcgoodfood.com.
type BBCGoodFoodScraper struct {
	schema *schema.RecipeScraper
}

func (m *BBCGoodFoodScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *BBCGoodFoodScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *BBCGoodFoodScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *BBCGoodFoodScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *BBCGoodFoodScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *BBCGoodFoodScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *BBCGoodFoodScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *BBCGoodFoodScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *BBCGoodFoodScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *BBCGoodFoodScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *BBCGoodFoodScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *BBCGoodFoodScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *BBCGoodFoodScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *BBCGoodFoodScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *BBCGoodFoodScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *BBCGoodFoodScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *BBCGoodFoodScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *BBCGoodFoodScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *BBCGoodFoodScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *BBCGoodFoodScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *BBCGoodFoodScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *BBCGoodFoodScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
