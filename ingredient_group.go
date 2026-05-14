package recipe

// IngredientGroup represents a group of ingredients with an optional purpose.
// For example, a recipe might have groups like "For the sauce" and "For the dough".
type IngredientGroup struct {
	// Purpose is the label for this ingredient group, e.g. "For the sauce".
	// May be empty for the default/ungrouped group.
	Purpose string
	// Ingredients are the ingredients in this group.
	Ingredients []string
}
