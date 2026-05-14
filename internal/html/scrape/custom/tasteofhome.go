package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const TasteOfHomeHost = "tasteofhome.com"

// NewTasteOfHomeScraper returns a new instance of TasteOfHomeScraper.
func NewTasteOfHomeScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &TasteOfHomeScraper{schema: s}, nil
}

// TasteOfHomeScraper is a custom recipe scraper for tasteofhome.com.
type TasteOfHomeScraper struct {
	schema *schema.RecipeScraper
}

func (m *TasteOfHomeScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *TasteOfHomeScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *TasteOfHomeScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *TasteOfHomeScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *TasteOfHomeScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *TasteOfHomeScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *TasteOfHomeScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *TasteOfHomeScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *TasteOfHomeScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *TasteOfHomeScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *TasteOfHomeScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *TasteOfHomeScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *TasteOfHomeScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *TasteOfHomeScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *TasteOfHomeScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *TasteOfHomeScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *TasteOfHomeScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *TasteOfHomeScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *TasteOfHomeScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *TasteOfHomeScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *TasteOfHomeScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *TasteOfHomeScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
