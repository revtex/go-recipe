package schema

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html"
	ld "github.com/kkyr/go-recipe/internal/html/scrape/schema/json-ld"

	"github.com/PuerkitoBio/goquery"
	"github.com/senseyeio/duration"
)

// NewRecipeScraper returns a new instance of RecipeScraper. The provided doc
// is scraped for a recipe and an error is returned if none is found.
func NewRecipeScraper(doc *goquery.Document) (*RecipeScraper, error) {
	node, err := ld.NewRecipeProcessor().GetRecipeNode(doc)
	if err != nil {
		return nil, fmt.Errorf("could not get recipe root from ld+json document: %w", err)
	}

	return newRecipeScraper(doc, node), nil
}

func newRecipeScraper(doc *goquery.Document, data map[string]any) *RecipeScraper {
	return &RecipeScraper{doc: doc, root: data}
}

var _ recipe.Scraper = (*RecipeScraper)(nil)

// RecipeScraper is a recipe scraper.
type RecipeScraper struct {
	doc  *goquery.Document
	root map[string]any
}

// Author is the author of the recipe.
func (r *RecipeScraper) Author() (string, bool) {
	if node, ok := r.root["author"].(map[string]any); ok {
		return getStringValue(node, "name")
	}

	return "", false
}

// Categories are the categories of the recipe, e.g. appetizer, entrée, etc.
func (r *RecipeScraper) Categories() ([]string, bool) {
	if s, ok := r.root["recipeCategory"].(string); ok {
		ss := strings.Split(s, ",")
		for i := range ss {
			ss[i] = html.CleanString(ss[i])
		}

		return ss, true
	}

	return getSliceValue(r.root, "recipeCategory")
}

// CookTime is the time it takes to actually cook the dish.
func (r *RecipeScraper) CookTime() (time.Duration, bool) {
	return getDurationValue(r.root, "cookTime")
}

// CookingMethod is the method of cooking used, e.g. frying, steaming, etc.
func (r *RecipeScraper) CookingMethod() (string, bool) {
	return getStringValue(r.root, "cookingMethod")
}

// Cuisine is the cuisine of the recipe, e.g. mexican-inspired, french, etc.
func (r *RecipeScraper) Cuisine() ([]string, bool) {
	return getSliceValue(r.root, "recipeCuisine")
}

// Description is the description of the recipe.
func (r *RecipeScraper) Description() (string, bool) {
	return getStringValue(r.root, "description")
}

// Equipment returns the equipment needed for the recipe.
// Checks schema.org tool property and WPRM (WordPress Recipe Maker) equipment.
func (r *RecipeScraper) Equipment() ([]string, bool) {
	// Check schema.org "tool" property
	if ss, ok := getSliceValue(r.root, "tool"); ok {
		return ss, true
	}

	// Check for WPRM equipment in the HTML
	if r.doc != nil {
		var equipment []string
		r.doc.Find(".wprm-recipe-equipment .wprm-recipe-equipment-name").Each(func(_ int, sel *goquery.Selection) {
			name := html.CleanString(sel.Text())
			if name != "" {
				equipment = append(equipment, name)
			}
		})
		if len(equipment) > 0 {
			return equipment, true
		}
	}

	return nil, false
}

// ImageURL is a URL to an image of the dish.
func (r *RecipeScraper) ImageURL() (string, bool) {
	node := r.root["image"]

	if s, ok := node.(string); ok {
		return html.CleanString(s), true
	}

	if m, ok := node.(map[string]any); ok {
		return getStringValue(m, "url")
	}

	if ss, ok := node.([]any); ok {
		for _, item := range ss {
			if s, ok := item.(string); ok && s != "" {
				return html.CleanString(s), true
			}

			if m, ok := item.(map[string]any); ok {
				if url, ok := getStringValue(m, "url"); ok {
					return url, true
				}
			}
		}
	}

	return "", false
}

// Ingredients are all the ingredients used in the recipe.
func (r *RecipeScraper) Ingredients() ([]string, bool) {
	if ss, ok := getSliceValue(r.root, "recipeIngredient"); ok {
		return ss, true
	}

	return getSliceValue(r.root, "ingredients")
}

// IngredientGroups returns ingredients organized into groups.
// Falls back to WPRM HTML groups, then wraps flat Ingredients() in a single group.
func (r *RecipeScraper) IngredientGroups() ([]recipe.IngredientGroup, bool) {
	// Try WPRM ingredient groups from HTML
	if r.doc != nil {
		var groups []recipe.IngredientGroup
		r.doc.Find(".wprm-recipe-ingredient-group").Each(func(_ int, sel *goquery.Selection) {
			var purpose string
			if header := sel.Find(".wprm-recipe-group-name"); header.Length() > 0 {
				purpose = html.CleanString(header.Text())
			}
			var ingredients []string
			sel.Find(".wprm-recipe-ingredient").Each(func(_ int, li *goquery.Selection) {
				ing := html.CleanString(li.Text())
				if ing != "" {
					ingredients = append(ingredients, ing)
				}
			})
			if len(ingredients) > 0 {
				groups = append(groups, recipe.IngredientGroup{
					Purpose:     purpose,
					Ingredients: ingredients,
				})
			}
		})
		if len(groups) > 0 {
			return groups, true
		}
	}

	// Check HowToSection-style ingredient groups in schema
	if sections, ok := r.root["recipeIngredient"].([]any); ok {
		var groups []recipe.IngredientGroup
		for _, section := range sections {
			if m, ok := section.(map[string]any); ok {
				if t, _ := m["type"].(string); t == "HowToSection" {
					purpose, _ := m["name"].(string)
					var ingredients []string
					if items, ok := m["itemListElement"].([]any); ok {
						for _, item := range items {
							if s, ok := item.(string); ok {
								ingredients = append(ingredients, html.CleanString(s))
							}
						}
					}
					if len(ingredients) > 0 {
						groups = append(groups, recipe.IngredientGroup{
							Purpose:     html.CleanString(purpose),
							Ingredients: ingredients,
						})
					}
				}
			}
		}
		if len(groups) > 0 {
			return groups, true
		}
	}

	// Fall back to wrapping flat ingredients in a single group
	if ingredients, ok := r.Ingredients(); ok {
		return []recipe.IngredientGroup{{Ingredients: ingredients}}, true
	}

	return nil, false
}

// Instructions are all the steps in making the recipe.
func (r *RecipeScraper) Instructions() ([]string, bool) {
	if nodes, ok := r.root["recipeInstructions"].([]any); ok {
		instructions := getInstructions(nodes)
		if len(instructions) > 0 {
			return instructions, true
		}
	}

	return nil, false
}

// Keywords returns the tags or keywords for the recipe.
func (r *RecipeScraper) Keywords() ([]string, bool) {
	if s, ok := r.root["keywords"].(string); ok {
		ss := strings.Split(s, ",")
		var keywords []string
		for _, k := range ss {
			k = html.CleanString(k)
			if k != "" {
				keywords = append(keywords, k)
			}
		}
		if len(keywords) > 0 {
			return keywords, true
		}
	}

	return getSliceValue(r.root, "keywords")
}

// Language is the language used in the recipe expressed in IETF BCP 47 standard.
func (r *RecipeScraper) Language() (string, bool) {
	for _, key := range []string{"inLanguage", "language"} {
		if s, ok := getStringValue(r.root, key); ok {
			return s, true
		}
	}

	return "", false
}

// Name is the name of the recipe.
func (r *RecipeScraper) Name() (string, bool) {
	return getStringValue(r.root, "name")
}

// Nutrition is nutritional information about the dish.
func (r *RecipeScraper) Nutrition() (recipe.Nutrition, bool) {
	if m, ok := r.root["nutrition"].(map[string]any); ok {
		return ParseNutritionalInformation(m), true
	}

	return recipe.Nutrition{}, false
}

// PrepTime is the length of time it takes to prepare the items to be used in the instructions.
func (r *RecipeScraper) PrepTime() (time.Duration, bool) {
	return getDurationValue(r.root, "prepTime")
}

// Ratings returns the average rating of the recipe from schema.org AggregateRating.
func (r *RecipeScraper) Ratings() (float32, bool) {
	if m := r.getAggregateRating(); m != nil {
		for _, key := range []string{"ratingValue", "value"} {
			if v, ok := parseFloatFromAny(m[key]); ok && v > 0 {
				return v, true
			}
		}
	}
	return 0, false
}

// RatingsCount returns the total number of ratings from schema.org AggregateRating.
func (r *RecipeScraper) RatingsCount() (int, bool) {
	if m := r.getAggregateRating(); m != nil {
		for _, key := range []string{"ratingCount", "reviewCount"} {
			if v, ok := parseIntFromAny(m[key]); ok && v > 0 {
				return v, true
			}
		}
	}
	return 0, false
}

// getAggregateRating returns the aggregateRating map from the recipe node,
// or falls back to scanning separate JSON-LD blocks that reference the same @id.
func (r *RecipeScraper) getAggregateRating() map[string]any {
	if m, ok := r.root["aggregateRating"].(map[string]any); ok {
		return m
	}

	// Some sites (e.g. BBC Good Food) put aggregateRating in a separate
	// JSON-LD block that references the Recipe via @id.
	if r.doc == nil {
		return nil
	}

	recipeID, _ := r.root["@id"].(string)
	if recipeID == "" {
		recipeID, _ = r.root["id"].(string)
	}
	if recipeID == "" {
		return nil
	}

	var result map[string]any

	r.doc.Find(`script[type="application/ld+json"]`).Each(func(_ int, sel *goquery.Selection) {
		if result != nil {
			return
		}

		var data map[string]any
		if err := json.Unmarshal([]byte(sel.Text()), &data); err != nil {
			return
		}

		id, _ := data["@id"].(string)
		if id != recipeID {
			return
		}

		if m, ok := data["aggregateRating"].(map[string]any); ok {
			result = m
		}
	})

	return result
}

// SiteName returns the name of the website. Checks OpenGraph og:site_name and the HTML title.
func (r *RecipeScraper) SiteName() (string, bool) {
	if r.doc != nil {
		// Try OpenGraph og:site_name (may use property= or name= attribute)
		sel := r.doc.Find(`meta[property="og:site_name"], meta[name="og:site_name"]`).First()
		if val, exists := sel.Attr("content"); exists {
			name := html.CleanString(val)
			if name != "" {
				return name, true
			}
		}
	}
	return "", false
}

// SuitableDiets indicates dietary restrictions or guidelines for which the recipe is suitable.
func (r *RecipeScraper) SuitableDiets() ([]recipe.Diet, bool) {
	if v, ok := getStringValue(r.root, "suitableForDiet"); ok {
		ss := strings.Split(v, ",")
		ret := make([]recipe.Diet, 0, len(ss))

		for _, s := range ss {
			if diet := ParseDiet(s); diet != recipe.UnknownDiet {
				ret = append(ret, diet)
			}
		}

		return ret, true
	}

	return nil, false
}

// TotalTime is the total time required to perform the instructions (including the prep time).
func (r *RecipeScraper) TotalTime() (time.Duration, bool) {
	return getDurationValue(r.root, "totalTime")
}

// Yields is the quantity that results from the recipe.
func (r *RecipeScraper) Yields() (string, bool) {
	yieldData := r.root["recipeYield"]
	if yieldData == nil {
		yieldData = r.root["yield"]
	}

	if s, ok := yieldData.(string); ok {
		return html.CleanString(s), true
	}

	if ss, ok := yieldData.([]any); ok {
		if len(ss) > 0 {
			// the last element seems to be the most descriptive
			if s, ok := ss[len(ss)-1].(string); ok {
				return html.CleanString(s), true
			}
		}
	}

	return "", false
}

func getInstructions(nodes []any) []string {
	var ret []string

	for _, instr := range nodes {
		if s, ok := instr.(string); ok {
			ret = append(ret, html.CleanString(s))
		}

		if m, ok := instr.(map[string]any); ok {
			if m["type"] == "HowToStep" {
				if step, ok := m["text"].(string); ok {
					ret = append(ret, html.CleanString(step))
				}
			} else if items, ok := m["itemListElement"].([]any); ok {
				ret = append(ret, getInstructions(items)...)
			}
		}
	}

	return ret
}

func getStringValue(node map[string]any, key string) (string, bool) {
	s, ok := node[key].(string)
	if !ok {
		return "", false
	}

	cleaned := html.CleanString(s)
	if cleaned == "" {
		return "", false
	}

	return cleaned, true
}

func getSliceValue(node map[string]any, key string) ([]string, bool) {
	if s, ok := getStringValue(node, key); ok {
		return []string{s}, true
	}

	if ss, ok := node[key].([]any); ok {
		ret := make([]string, 0, len(ss))

		for _, v := range ss {
			if s, ok := v.(string); ok {
				ret = append(ret, html.CleanString(s))
			}
		}

		return ret, true
	}

	return nil, false
}

func getDurationValue(node map[string]any, key string) (time.Duration, bool) {
	v, ok := node[key].(string)
	if !ok {
		return 0, false
	}

	td, ok := parseDuration(v)
	if !ok {
		return 0, false
	}

	if td == 0 {
		return 0, false
	}

	return td, true
}

// looseDurationRE matches ISO 8601-ish duration components.
// Used as a fallback for malformed values like "PT0D0H30M" (which puts
// D after T) that some recipe sites emit.
var looseDurationRE = regexp.MustCompile(`(\d+(?:\.\d+)?)\s*([DHMS])`)

// parseDuration parses an ISO 8601 duration, falling back to a tolerant
// regex-based parser for malformed values that real-world recipe sites
// occasionally produce.
func parseDuration(v string) (time.Duration, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0, false
	}

	if dur, err := duration.ParseISO8601(v); err == nil {
		var td time.Duration
		td += time.Duration(dur.TH) * time.Hour
		td += time.Duration(dur.TM) * time.Minute
		td += time.Duration(dur.TS) * time.Second

		if td > 0 {
			return td, true
		}
	}

	// Fallback: pull all <num><unit> pairs out of the string. This handles
	// malformed values like "PT0D0H30M" (D must come before T per spec).
	matches := looseDurationRE.FindAllStringSubmatch(strings.ToUpper(v), -1)
	if len(matches) == 0 {
		return 0, false
	}

	var td time.Duration

	for _, m := range matches {
		n, err := strconv.ParseFloat(m[1], 64)
		if err != nil {
			continue
		}

		switch m[2] {
		case "D":
			td += time.Duration(n * float64(24*time.Hour))
		case "H":
			td += time.Duration(n * float64(time.Hour))
		case "M":
			td += time.Duration(n * float64(time.Minute))
		case "S":
			td += time.Duration(n * float64(time.Second))
		}
	}

	return td, td > 0
}

func parseFloatFromAny(v any) (float32, bool) {
	switch val := v.(type) {
	case float64:
		return float32(val), true
	case float32:
		return val, true
	case string:
		f, err := strconv.ParseFloat(strings.TrimSpace(val), 32)
		if err != nil {
			return 0, false
		}
		return float32(f), true
	}
	return 0, false
}

func parseIntFromAny(v any) (int, bool) {
	switch val := v.(type) {
	case float64:
		return int(val), true
	case string:
		i, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return 0, false
		}
		return i, true
	}
	return 0, false
}
