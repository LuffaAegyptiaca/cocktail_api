package models

import "database/sql"

type CocktailIngredient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewCocktailIngredient(db *sql.DB) *CocktailIngredient {
	return &CocktailIngredient{}
}