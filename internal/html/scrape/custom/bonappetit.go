package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const BonAppetitHost = "bonappetit.com"

// NewBonAppetitScraper returns a new instance of BonAppetitScraper.
func NewBonAppetitScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &BonAppetitScraper{schema: s}, nil
}

// BonAppetitScraper is a custom recipe scraper for bonappetit.com.
type BonAppetitScraper struct {
	schema *schema.RecipeScraper
}

func (m *BonAppetitScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *BonAppetitScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *BonAppetitScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *BonAppetitScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *BonAppetitScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *BonAppetitScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *BonAppetitScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *BonAppetitScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *BonAppetitScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *BonAppetitScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *BonAppetitScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *BonAppetitScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *BonAppetitScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *BonAppetitScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *BonAppetitScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *BonAppetitScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *BonAppetitScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *BonAppetitScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *BonAppetitScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *BonAppetitScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *BonAppetitScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *BonAppetitScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
