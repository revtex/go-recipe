package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const BudgetBytesHost = "budgetbytes.com"

// NewBudgetBytesScraper returns a new instance of BudgetBytesScraper.
func NewBudgetBytesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &BudgetBytesScraper{schema: s}, nil
}

// BudgetBytesScraper is a custom recipe scraper for budgetbytes.com.
type BudgetBytesScraper struct {
	schema *schema.RecipeScraper
}

func (m *BudgetBytesScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *BudgetBytesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *BudgetBytesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *BudgetBytesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *BudgetBytesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *BudgetBytesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *BudgetBytesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *BudgetBytesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *BudgetBytesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *BudgetBytesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *BudgetBytesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *BudgetBytesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *BudgetBytesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *BudgetBytesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *BudgetBytesScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *BudgetBytesScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *BudgetBytesScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *BudgetBytesScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *BudgetBytesScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *BudgetBytesScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *BudgetBytesScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *BudgetBytesScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
