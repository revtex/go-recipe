package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const SkinnyTasteHost = "skinnytaste.com"

// NewSkinnyTasteScraper returns a new instance of SkinnyTasteScraper.
func NewSkinnyTasteScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &SkinnyTasteScraper{schema: s}, nil
}

// SkinnyTasteScraper is a custom recipe scraper for skinnytaste.com.
type SkinnyTasteScraper struct {
	schema *schema.RecipeScraper
}

func (m *SkinnyTasteScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *SkinnyTasteScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *SkinnyTasteScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *SkinnyTasteScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *SkinnyTasteScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *SkinnyTasteScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *SkinnyTasteScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *SkinnyTasteScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *SkinnyTasteScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *SkinnyTasteScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *SkinnyTasteScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *SkinnyTasteScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *SkinnyTasteScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *SkinnyTasteScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *SkinnyTasteScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *SkinnyTasteScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *SkinnyTasteScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *SkinnyTasteScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *SkinnyTasteScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *SkinnyTasteScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *SkinnyTasteScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *SkinnyTasteScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
