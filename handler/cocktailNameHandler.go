package handler

import (
	"cocktail_api/database"
	"cocktail_api/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCocktailsIngredient(c echo.Context) error {
	db := database.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM cocktail_recipe")

	if err != nil {
		fmt.Println("err: ", err)
	}
	var cim_l []models.CocktailRecipes
	for rows.Next() {
		var cim models.CocktailRecipes
		rows.Scan(&cim.Id, &cim.Name, &cim.Recipe, &cim.CocktailStyle, &cim.Alcohol)
		cim_l = append(cim_l, cim)
	}
	fmt.Printf("%v", cim_l)
	// テスト中なので、適当な感じのものを返す
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func ReturnHoge() string {
	return "hgoe"
}
