package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const MarthaStewartHost = "marthastewart.com"

// NewMarthaStewartScraper returns a new instance of MarthaStewartScraper.
func NewMarthaStewartScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &MarthaStewartScraper{schema: s}, nil
}

// MarthaStewartScraper is a custom recipe scraper for marthastewart.com.
type MarthaStewartScraper struct {
	schema *schema.RecipeScraper
}

func (m *MarthaStewartScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *MarthaStewartScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *MarthaStewartScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *MarthaStewartScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *MarthaStewartScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *MarthaStewartScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *MarthaStewartScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *MarthaStewartScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *MarthaStewartScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *MarthaStewartScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *MarthaStewartScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *MarthaStewartScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *MarthaStewartScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *MarthaStewartScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *MarthaStewartScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *MarthaStewartScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *MarthaStewartScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *MarthaStewartScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *MarthaStewartScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *MarthaStewartScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *MarthaStewartScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *MarthaStewartScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
