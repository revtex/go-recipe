package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const TheKitchnHost = "thekitchn.com"

// NewTheKitchnScraper returns a new instance of TheKitchnScraper.
func NewTheKitchnScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &TheKitchnScraper{schema: s}, nil
}

// TheKitchnScraper is a custom recipe scraper for thekitchn.com.
type TheKitchnScraper struct {
	schema *schema.RecipeScraper
}

func (m *TheKitchnScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *TheKitchnScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *TheKitchnScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *TheKitchnScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *TheKitchnScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *TheKitchnScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *TheKitchnScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *TheKitchnScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *TheKitchnScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *TheKitchnScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *TheKitchnScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *TheKitchnScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *TheKitchnScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *TheKitchnScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *TheKitchnScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *TheKitchnScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *TheKitchnScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *TheKitchnScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *TheKitchnScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *TheKitchnScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *TheKitchnScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *TheKitchnScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
