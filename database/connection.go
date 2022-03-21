package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConf struct {
	Host     string
	Port     string
	User     string
	Password string
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
func Connect() *sql.DB {
	df, err := getDbEnv()
	if err != nil {
		fmt.Println("getDbEnv err")
	}
	dbName := "cocktail"
	db, err := sql.Open("postgres", "host="+df.Host+" user="+df.User+" password="+df.Password+" dbname="+dbName+" port="+df.Port+" sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db
}
