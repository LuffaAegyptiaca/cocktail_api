package cocktail

import "database/sql"

type CocktailRecipes struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Recipe        string `json:"recipe"`
	CocktailStyle string `json:"cocktail_style"`
	Alcohol       int    `json:"alcohol"`
}

func NewCocktailRecipes(db *sql.DB) *CocktailRecipes {
	return &CocktailRecipes{}
}
