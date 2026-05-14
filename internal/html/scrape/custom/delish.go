package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const DelishHost = "delish.com"

// NewDelishScraper returns a new instance of DelishScraper.
func NewDelishScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &DelishScraper{schema: s}, nil
}

// DelishScraper is a custom recipe scraper for delish.com.
type DelishScraper struct {
	schema *schema.RecipeScraper
}

func (m *DelishScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *DelishScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *DelishScraper) CookTime() (time.Duration, bool) {
	// Delish always sets cookTime to PT0S in their schema.
	// Calculate cook time from total time minus prep time instead.
	total, totalOK := m.schema.TotalTime()
	prep, prepOK := m.schema.PrepTime()
	if totalOK && prepOK && total > prep {
		return total - prep, true
	}
	return 0, false
}

func (m *DelishScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *DelishScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *DelishScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *DelishScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *DelishScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *DelishScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *DelishScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *DelishScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *DelishScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *DelishScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *DelishScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *DelishScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *DelishScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *DelishScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *DelishScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *DelishScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *DelishScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *DelishScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *DelishScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
