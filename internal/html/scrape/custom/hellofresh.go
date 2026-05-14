package custom

import (
	"fmt"
	"strings"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const HelloFreshHost = "hellofresh.com"

// NewHelloFreshScraper returns a new instance of HelloFreshScraper.
func NewHelloFreshScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &HelloFreshScraper{schema: s}, nil
}

// HelloFreshScraper is a custom recipe scraper for hellofresh.com.
type HelloFreshScraper struct {
	schema *schema.RecipeScraper
}

func (m *HelloFreshScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *HelloFreshScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *HelloFreshScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *HelloFreshScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *HelloFreshScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *HelloFreshScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *HelloFreshScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

// Instructions returns the recipe instructions. hellofresh.com stores every
// step as a single HowToStep whose text is an HTML fragment containing a
// <ul><li>...</li></ul> list of sub-steps, so the default schema parser
// returns each step as a wall of HTML markup. Override it to split on the
// list items and strip the tags, producing one entry per actual sub-step.
func (m *HelloFreshScraper) Instructions() ([]string, bool) {
	raw, ok := m.schema.Instructions()
	if !ok {
		return nil, false
	}

	var steps []string
	for _, r := range raw {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(r))
		if err != nil {
			steps = append(steps, html.CleanString(r))
			continue
		}
		items := doc.Find("li")
		if items.Length() == 0 {
			steps = append(steps, html.CleanString(doc.Text()))
			continue
		}
		items.Each(func(_ int, sel *goquery.Selection) {
			if s := html.CleanString(sel.Text()); s != "" {
				steps = append(steps, s)
			}
		})
	}

	if len(steps) == 0 {
		return nil, false
	}
	return steps, true
}

func (m *HelloFreshScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *HelloFreshScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *HelloFreshScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *HelloFreshScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *HelloFreshScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *HelloFreshScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *HelloFreshScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *HelloFreshScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *HelloFreshScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *HelloFreshScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *HelloFreshScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *HelloFreshScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *HelloFreshScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *HelloFreshScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
