package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const SeriousEatsHost = "seriouseats.com"

// NewSeriousEatsScraper returns a new instance of SeriousEatsScraper.
func NewSeriousEatsScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &SeriousEatsScraper{schema: s}, nil
}

// SeriousEatsScraper is a custom recipe scraper for seriouseats.com.
type SeriousEatsScraper struct {
	schema *schema.RecipeScraper
}

func (m *SeriousEatsScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *SeriousEatsScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *SeriousEatsScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *SeriousEatsScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *SeriousEatsScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *SeriousEatsScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *SeriousEatsScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *SeriousEatsScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *SeriousEatsScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *SeriousEatsScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *SeriousEatsScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *SeriousEatsScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *SeriousEatsScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *SeriousEatsScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *SeriousEatsScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *SeriousEatsScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *SeriousEatsScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *SeriousEatsScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *SeriousEatsScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *SeriousEatsScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *SeriousEatsScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *SeriousEatsScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
