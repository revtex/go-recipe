package custom

import (
	"fmt"
	"regexp"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const HalfBakedHarvestHost = "halfbakedharvest.com"

// stepNumberRE matches a leading step number such as "1. " or "12) ".
var stepNumberRE = regexp.MustCompile(`^\d+[.)]\s*`)

// NewHalfBakedHarvestScraper returns a new instance of HalfBakedHarvestScraper.
func NewHalfBakedHarvestScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &HalfBakedHarvestScraper{doc: doc, schema: s}, nil
}

// HalfBakedHarvestScraper is a custom recipe scraper for halfbakedharvest.com.
type HalfBakedHarvestScraper struct {
	doc    *goquery.Document
	schema *schema.RecipeScraper
}

func (m *HalfBakedHarvestScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *HalfBakedHarvestScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *HalfBakedHarvestScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *HalfBakedHarvestScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *HalfBakedHarvestScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *HalfBakedHarvestScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *HalfBakedHarvestScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

// Instructions returns the recipe instructions. halfbakedharvest.com (powered
// by WP Recipe Maker) packs every numbered step into a single HowToStep in its
// JSON-LD output, so we instead extract one step per <span> inside the
// .wprm-recipe-instruction-text element.
func (m *HalfBakedHarvestScraper) Instructions() ([]string, bool) {
	var steps []string
	m.doc.Find(".wprm-recipe-instruction-text > span").Each(func(_ int, sel *goquery.Selection) {
		s := html.CleanString(sel.Text())
		s = stepNumberRE.ReplaceAllString(s, "")
		if s != "" {
			steps = append(steps, s)
		}
	})
	if len(steps) > 0 {
		return steps, true
	}
	return m.schema.Instructions()
}

func (m *HalfBakedHarvestScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *HalfBakedHarvestScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *HalfBakedHarvestScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *HalfBakedHarvestScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *HalfBakedHarvestScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *HalfBakedHarvestScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *HalfBakedHarvestScraper) Yields() (string, bool) {
	return m.schema.Yields()
}

func (m *HalfBakedHarvestScraper) CookingMethod() (string, bool) {
	return m.schema.CookingMethod()
}

func (m *HalfBakedHarvestScraper) Equipment() ([]string, bool) {
	return m.schema.Equipment()
}

func (m *HalfBakedHarvestScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	return m.schema.IngredientGroups()
}

func (m *HalfBakedHarvestScraper) Keywords() ([]string, bool) {
	return m.schema.Keywords()
}

func (m *HalfBakedHarvestScraper) Ratings() (float32, bool) {
	return m.schema.Ratings()
}

func (m *HalfBakedHarvestScraper) RatingsCount() (int, bool) {
	return m.schema.RatingsCount()
}

func (m *HalfBakedHarvestScraper) SiteName() (string, bool) {
	return m.schema.SiteName()
}
