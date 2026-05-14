package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const TastyHost = "tasty.co"

// NewTastyScraper returns a new instance of TastyScraper.
func NewTastyScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &TastyScraper{schema: s}, nil
}

// TastyScraper is a custom recipe scraper for tasty.co.
type TastyScraper struct {
	schema *schema.RecipeScraper
}

func (m *TastyScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *TastyScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *TastyScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *TastyScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *TastyScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *TastyScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *TastyScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *TastyScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *TastyScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *TastyScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *TastyScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *TastyScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *TastyScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *TastyScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *TastyScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *TastyScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *TastyScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *TastyScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *TastyScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *TastyScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *TastyScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *TastyScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
