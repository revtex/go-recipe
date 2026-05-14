package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const NYTimesHost = "cooking.nytimes.com"

// NewNYTimesScraper returns a new instance of NYTimesScraper.
func NewNYTimesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &NYTimesScraper{schema: s}, nil
}

// NYTimesScraper is a custom recipe scraper for cooking.nytimes.com.
type NYTimesScraper struct {
	schema *schema.RecipeScraper
}

func (m *NYTimesScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *NYTimesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *NYTimesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *NYTimesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *NYTimesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *NYTimesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *NYTimesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *NYTimesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *NYTimesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *NYTimesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *NYTimesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *NYTimesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *NYTimesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *NYTimesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *NYTimesScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *NYTimesScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *NYTimesScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *NYTimesScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *NYTimesScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *NYTimesScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *NYTimesScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *NYTimesScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
