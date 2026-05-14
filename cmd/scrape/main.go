package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kkyr/go-recipe"
	pkg "github.com/kkyr/go-recipe/pkg/recipe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("Scraping: %s\n\n", url)

	scraper, err := pkg.ScrapeURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if v, ok := scraper.Name(); ok {
		fmt.Printf("Name:          %s\n", v)
	}
	if v, ok := scraper.Author(); ok {
		fmt.Printf("Author:        %s\n", v)
	}
	if v, ok := scraper.SiteName(); ok {
		fmt.Printf("SiteName:      %s\n", v)
	}
	if v, ok := scraper.Description(); ok {
		fmt.Printf("Description:   %.100s...\n", v)
	}
	if v, ok := scraper.ImageURL(); ok {
		fmt.Printf("ImageURL:      %s\n", v)
	}
	if v, ok := scraper.Language(); ok {
		fmt.Printf("Language:      %s\n", v)
	}
	if v, ok := scraper.Yields(); ok {
		fmt.Printf("Yields:        %s\n", v)
	}
	if v, ok := scraper.PrepTime(); ok {
		fmt.Printf("PrepTime:      %s\n", v)
	}
	if v, ok := scraper.CookTime(); ok {
		fmt.Printf("CookTime:      %s\n", v)
	}
	if v, ok := scraper.TotalTime(); ok {
		fmt.Printf("TotalTime:     %s\n", v)
	}
	if v, ok := scraper.CookingMethod(); ok {
		fmt.Printf("CookingMethod: %s\n", v)
	}
	if v, ok := scraper.Categories(); ok {
		fmt.Printf("Categories:    %v\n", v)
	}
	if v, ok := scraper.Cuisine(); ok {
		fmt.Printf("Cuisine:       %v\n", v)
	}
	if v, ok := scraper.Keywords(); ok {
		fmt.Printf("Keywords:      %v\n", v)
	}
	if v, ok := scraper.Ratings(); ok {
		fmt.Printf("Ratings:       %.1f\n", v)
	}
	if v, ok := scraper.RatingsCount(); ok {
		fmt.Printf("RatingsCount:  %d\n", v)
	}
	if v, ok := scraper.SuitableDiets(); ok {
		fmt.Printf("Diets:         %v\n", v)
	}
	if v, ok := scraper.Equipment(); ok {
		fmt.Printf("Equipment:     %v\n", v)
	}
	if v, ok := scraper.Nutrition(); ok {
		fmt.Printf("Nutrition:     Cal:%.0f Carb:%.0fg Fat:%.0fg Protein:%.0fg\n",
			v.Calories, v.CarbohydrateGrams, v.FatGrams, v.ProteinGrams)
	}

	fmt.Println()
	if v, ok := scraper.Ingredients(); ok {
		fmt.Printf("Ingredients (%d):\n", len(v))
		for _, ing := range v {
			fmt.Printf("  - %s\n", ing)
		}
	}
	if v, ok := scraper.IngredientGroups(); ok {
		if len(v) > 1 || (len(v) == 1 && v[0].Purpose != "") {
			fmt.Printf("\nIngredient Groups (%d):\n", len(v))
			for _, g := range v {
				if g.Purpose != "" {
					fmt.Printf("  [%s]\n", g.Purpose)
				} else {
					fmt.Printf("  [ungrouped]\n")
				}
				for _, ing := range g.Ingredients {
					fmt.Printf("    - %s\n", ing)
				}
			}
		}
	}

	fmt.Println()
	if v, ok := scraper.Instructions(); ok {
		fmt.Printf("Instructions (%d):\n", len(v))
		for i, step := range v {
			fmt.Printf("  %d. %.120s\n", i+1, step)
		}
	}

	fmt.Printf("\n--- JSON ---\n")
	data, err := recipe.ToJSON(scraper)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON error: %v\n", err)
		os.Exit(1)
	}
	var pretty json.RawMessage = data
	out, _ := json.MarshalIndent(pretty, "", "  ")
	fmt.Println(string(out))
}
