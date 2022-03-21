package models

type CocktailIngredientQuantity struct {
	CocktailRecipeId     int    `json:"cocktail_recipe_id"`
	CocktailIngredientId int    `json:"cocktail_ingredient_id"`
	Quantity             int    `json:"quantity"`
	Unit                 string `json:"unit"`
	BaseFlag             bool   `json:"base_flag"`
}
