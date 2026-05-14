package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const SimplyRecipesHost = "simplyrecipes.com"

// NewSimplyRecipesScraper returns a new instance of SimplyRecipesScraper.
func NewSimplyRecipesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &SimplyRecipesScraper{schema: s}, nil
}

// SimplyRecipesScraper is a custom recipe scraper for simplyrecipes.com.
type SimplyRecipesScraper struct {
	schema *schema.RecipeScraper
}

func (m *SimplyRecipesScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *SimplyRecipesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *SimplyRecipesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *SimplyRecipesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *SimplyRecipesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *SimplyRecipesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *SimplyRecipesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *SimplyRecipesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *SimplyRecipesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *SimplyRecipesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *SimplyRecipesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *SimplyRecipesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *SimplyRecipesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *SimplyRecipesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *SimplyRecipesScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *SimplyRecipesScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *SimplyRecipesScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *SimplyRecipesScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *SimplyRecipesScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *SimplyRecipesScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *SimplyRecipesScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *SimplyRecipesScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
