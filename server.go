package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type DBConf struct {
	Host     string
	Port     string
	User     string
	Password string
}

func main() {
	fmt.Println("start echo server")
	// Echo instance
	e := echo.New()
	// db_stady()
	df, err := getDbEnv()
	if err != nil {
		println(err)
	}
	fmt.Println(df)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// envからDB関係の環境変数を取得する関数
func getDbEnv() (df *DBConf, err error) {
	// 環境変数を読み込み
	err = godotenv.Load()

	// エラーチェック
	if err != nil {
		return nil, err
	}

	// DBの各種変数を環境変数から取得
	df = &DBConf{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	}
	// チェック
	// fmt.Printf("HOST = %s, PORT = %s, USER = %s, PASSWORD = %s", df.Host, df.Port, df.User, df.Password)
	// print(df)
	return df, nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// gormを使ってdbの接続をする関数
// SQL+プリペアードステートメントを利用する予定なので、最終的には削除する予定
func db_stady() {
	dsn := "host=localhost user=test password=test dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
	}
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user)
	println("[stmn] db status")
	println(user.ID)
	println(result.Error)
	println((result.RowsAffected))
	// result = db.Find(&user).Row()
	println(result.Scan(&user))
}
