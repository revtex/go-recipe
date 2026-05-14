package recipe

import (
	"encoding/json"
	"time"
)

// Scraper is a type that is returns recipe information from an underlying data source.
type Scraper interface {
	// Author is the author of the recipe.
	Author() (string, bool)
	// Categories are the categories of the recipe, e.g. appetizer, entrée, etc.
	Categories() ([]string, bool)
	// CookTime is the time it takes to actually cook the dish.
	CookTime() (time.Duration, bool)
	// CookingMethod is the method of cooking used, e.g. frying, steaming, etc.
	CookingMethod() (string, bool)
	// Cuisine is the cuisine of the recipe, e.g. mexican-inspired, french, etc.
	Cuisine() ([]string, bool)
	// Description is the description of the recipe.
	Description() (string, bool)
	// Equipment returns the equipment needed for the recipe.
	Equipment() ([]string, bool)
	// ImageURL is a URL to an image of the dish.
	ImageURL() (string, bool)
	// IngredientGroups returns the ingredients organized into groups with optional labels.
	IngredientGroups() ([]IngredientGroup, bool)
	// Ingredients are all the ingredients used in the recipe.
	Ingredients() ([]string, bool)
	// Instructions are all the steps in making the recipe.
	Instructions() ([]string, bool)
	// Keywords returns the tags or keywords used to describe the recipe.
	Keywords() ([]string, bool)
	// Language is the language used in the recipe expressed in IETF BCP 47 standard.
	Language() (string, bool)
	// Name is the name of the recipe.
	Name() (string, bool)
	// Nutrition is nutritional information about the dish.
	Nutrition() (Nutrition, bool)
	// PrepTime is the length of time it takes to prepare the items to be used in the instructions.
	PrepTime() (time.Duration, bool)
	// Ratings returns the average rating of the recipe.
	Ratings() (float32, bool)
	// RatingsCount returns the total number of ratings of the recipe.
	RatingsCount() (int, bool)
	// SiteName returns the name of the website the recipe is from.
	SiteName() (string, bool)
	// SuitableDiets indicates dietary restrictions or guidelines for which the recipe is suitable.
	SuitableDiets() ([]Diet, bool)
	// TotalTime is the total time required to perform the instructions (including the prep time).
	TotalTime() (time.Duration, bool)
	// Yields is the quantity that results from the recipe.
	Yields() (string, bool)
}

// ToJSON serializes the recipe scraper's data to a JSON byte slice.
// Fields that are not available (ok == false) are omitted.
func ToJSON(s Scraper) ([]byte, error) {
	m := make(map[string]any)

	if v, ok := s.Author(); ok {
		m["author"] = v
	}
	if v, ok := s.Categories(); ok {
		m["categories"] = v
	}
	if v, ok := s.CookTime(); ok {
		m["cookTime"] = v.String()
	}
	if v, ok := s.CookingMethod(); ok {
		m["cookingMethod"] = v
	}
	if v, ok := s.Cuisine(); ok {
		m["cuisine"] = v
	}
	if v, ok := s.Description(); ok {
		m["description"] = v
	}
	if v, ok := s.Equipment(); ok {
		m["equipment"] = v
	}
	if v, ok := s.ImageURL(); ok {
		m["imageURL"] = v
	}
	if v, ok := s.IngredientGroups(); ok {
		m["ingredientGroups"] = v
	}
	if v, ok := s.Ingredients(); ok {
		m["ingredients"] = v
	}
	if v, ok := s.Instructions(); ok {
		m["instructions"] = v
	}
	if v, ok := s.Keywords(); ok {
		m["keywords"] = v
	}
	if v, ok := s.Language(); ok {
		m["language"] = v
	}
	if v, ok := s.Name(); ok {
		m["name"] = v
	}
	if v, ok := s.Nutrition(); ok {
		m["nutrition"] = v
	}
	if v, ok := s.PrepTime(); ok {
		m["prepTime"] = v.String()
	}
	if v, ok := s.Ratings(); ok {
		m["ratings"] = v
	}
	if v, ok := s.RatingsCount(); ok {
		m["ratingsCount"] = v
	}
	if v, ok := s.SiteName(); ok {
		m["siteName"] = v
	}
	if v, ok := s.SuitableDiets(); ok {
		m["suitableDiets"] = v
	}
	if v, ok := s.TotalTime(); ok {
		m["totalTime"] = v.String()
	}
	if v, ok := s.Yields(); ok {
		m["yields"] = v
	}

	return json.Marshal(m)
}
