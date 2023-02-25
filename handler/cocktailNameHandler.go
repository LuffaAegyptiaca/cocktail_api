package handler

import (
	"cocktail_api/database"
	"cocktail_api/models/cocktail"
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

	resData := map[string][]cocktail.CocktailRecipes{}
	resData["cocktail_recipes"] = nil
	// データ読み込み
	for rows.Next() {
		var cim cocktail.CocktailRecipes
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

// TODO:関数名を変える必要あり。
func GetCocktailsIngredientBase(c echo.Context) error {
	// DB接続~クエリ実行まで
	db := database.Connect()
	defer db.Close()
	rows, err := db.Query(`
		SELECT
			cr.id,
			cr.name,
			cr.recipe,
			cr.cocktail_style,
			cr.alcohol
		FROM
			cocktail_recipe AS cr
		INNER JOIN cocktail_ingredient_quantity AS ciq ON ciq.cocktail_recipe_id = cr.id
		INNER JOIN cocktail_ingredient          AS ci  ON ciq.cocktail_ingredient_id = ci.id
		WHERE
			ciq.base_flag = true
			AND ci.name = $1`, c.Param("name"))
	if err != nil {
		fmt.Println("err: ", err)
	}

	resData := map[string][]cocktail.CocktailRecipes{}
	resData["cocktail_recipes"] = nil
	// データ読み込み
	for rows.Next() {
		var cim cocktail.CocktailRecipes
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
