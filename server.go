package main

import (
	"cocktail_api/handler"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	fmt.Println("start echo server")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	// 起動確認用。最終的に消す。
	e.GET("/", hello)

	// 全レシピ返すやつ
	// 件数多くない＆追加はしない仕様なので、ページングはAPP側に任せる
	// クエリパラメータは共通のパラメータを利用する。
	// 項目はアルコール度数(under-alcohol, over-alcohol)と飲み方（type（ショートカクテルとか））
	e.GET("/v1/cocktails", handler.CocktailAllRecipes)
	e.GET("/v1/cocktails/:base", handler.CocktailRecipes)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
