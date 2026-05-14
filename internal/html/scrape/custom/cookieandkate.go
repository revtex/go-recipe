package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const CookieAndKateHost = "cookieandkate.com"

// NewCookieAndKateScraper returns a new instance of CookieAndKateScraper.
func NewCookieAndKateScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &CookieAndKateScraper{schema: s}, nil
}

// CookieAndKateScraper is a custom recipe scraper for cookieandkate.com.
type CookieAndKateScraper struct {
	schema *schema.RecipeScraper
}

func (m *CookieAndKateScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *CookieAndKateScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *CookieAndKateScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *CookieAndKateScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *CookieAndKateScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *CookieAndKateScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *CookieAndKateScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *CookieAndKateScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *CookieAndKateScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *CookieAndKateScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *CookieAndKateScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *CookieAndKateScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *CookieAndKateScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *CookieAndKateScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *CookieAndKateScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *CookieAndKateScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *CookieAndKateScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *CookieAndKateScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *CookieAndKateScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *CookieAndKateScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *CookieAndKateScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *CookieAndKateScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
