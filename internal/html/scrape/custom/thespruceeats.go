package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const TheSpruceEatsHost = "thespruceeats.com"

// NewTheSpruceEatsScraper returns a new instance of TheSpruceEatsScraper.
func NewTheSpruceEatsScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &TheSpruceEatsScraper{schema: s}, nil
}

// TheSpruceEatsScraper is a custom recipe scraper for thespruceeats.com.
type TheSpruceEatsScraper struct {
	schema *schema.RecipeScraper
}

func (m *TheSpruceEatsScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *TheSpruceEatsScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *TheSpruceEatsScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *TheSpruceEatsScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *TheSpruceEatsScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *TheSpruceEatsScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *TheSpruceEatsScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *TheSpruceEatsScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *TheSpruceEatsScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *TheSpruceEatsScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *TheSpruceEatsScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *TheSpruceEatsScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *TheSpruceEatsScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *TheSpruceEatsScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *TheSpruceEatsScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *TheSpruceEatsScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *TheSpruceEatsScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *TheSpruceEatsScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *TheSpruceEatsScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *TheSpruceEatsScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *TheSpruceEatsScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *TheSpruceEatsScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
