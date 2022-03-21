package handler

import (
	"cocktail_api/database"
	"cocktail_api/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCocktailsIngredient(c echo.Context) error {
	// DB接続~クエリ実行まで
	db := database.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM cocktail_recipe")
	if err != nil {
		fmt.Println("err: ", err)
	}

	resData := map[string][]models.CocktailRecipes{}
	resData["cocktail_recipes"] = nil
	// データ読み込み
	for rows.Next() {
		var cim models.CocktailRecipes
		rows.Scan(&cim.Id, &cim.Name, &cim.Recipe, &cim.CocktailStyle, &cim.Alcohol)
		resData["cocktail_recipes"] = append(resData["cocktail_recipes"], cim)
	}
	// デバック用
	fmt.Printf("%v", resData)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, resData)
}

func ReturnHoge() string {
	return "hgoe"
}