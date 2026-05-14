package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const PinchOfYumHost = "pinchofyum.com"

// NewPinchOfYumScraper returns a new instance of PinchOfYumScraper.
func NewPinchOfYumScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &PinchOfYumScraper{schema: s}, nil
}

// PinchOfYumScraper is a custom recipe scraper for pinchofyum.com.
type PinchOfYumScraper struct {
	schema *schema.RecipeScraper
}

func (m *PinchOfYumScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *PinchOfYumScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *PinchOfYumScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *PinchOfYumScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *PinchOfYumScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *PinchOfYumScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *PinchOfYumScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *PinchOfYumScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *PinchOfYumScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *PinchOfYumScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *PinchOfYumScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *PinchOfYumScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *PinchOfYumScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *PinchOfYumScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *PinchOfYumScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *PinchOfYumScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *PinchOfYumScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *PinchOfYumScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *PinchOfYumScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *PinchOfYumScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *PinchOfYumScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *PinchOfYumScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
