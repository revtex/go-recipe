package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const EpicuriousHost = "epicurious.com"

// NewEpicuriousScraper returns a new instance of EpicuriousScraper.
func NewEpicuriousScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &EpicuriousScraper{doc: doc, schema: s}, nil
}

// EpicuriousScraper is a custom recipe scraper for epicurious.com.
type EpicuriousScraper struct {
	doc    *goquery.Document
	schema *schema.RecipeScraper
}

func (m *EpicuriousScraper) Author() (string, bool) {
	// Epicurious stores the author in an itemprop="author" anchor tag.
	ss := m.doc.Find("a[itemprop='author']").Map(func(_ int, sel *goquery.Selection) string {
		return sel.Text()
	})
	if len(ss) > 0 {
		return html.CleanString(ss[0]), true
	}

	return m.schema.Author()
}

func (m *EpicuriousScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *EpicuriousScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *EpicuriousScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *EpicuriousScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *EpicuriousScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *EpicuriousScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *EpicuriousScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *EpicuriousScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *EpicuriousScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *EpicuriousScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *EpicuriousScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *EpicuriousScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *EpicuriousScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *EpicuriousScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *EpicuriousScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *EpicuriousScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *EpicuriousScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *EpicuriousScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *EpicuriousScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *EpicuriousScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *EpicuriousScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
