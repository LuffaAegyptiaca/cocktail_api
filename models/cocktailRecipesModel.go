package models

type CocktailRecipes struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Recipe        string `json:"recipe"`
	CocktailStyle string `json:"cocktail_style"`
	Alcohol       int    `json:"alcohol"`
}
