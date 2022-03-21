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
	Database string
}

// envからDB関係の環境変数を取得する関数
func getDbEnv() (df *DBConf, err error) {
	// 環境変数を定義したファイルから読み込み
	// TODO: 開発途中なので、環境変数を固定値にしているが、開発用と本番用の切り替えをできるようにする
	err = godotenv.Load(".env")

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
		Database: os.Getenv("DATABASE"),
	}
	// チェック用。基本コメント化。
	// TODO: デバックフラグを立てているときは、出力できるようにするなどして、確認したい時に都度コメントを外す運用をなくす。
	// fmt.Printf("HOST = %s, PORT = %s, USER = %s, PASSWORD = %s", df.Host, df.Port, df.User, df.Password)
	return df, nil
}
func Connect() *sql.DB {
	df, err := getDbEnv()
	if err != nil {
		fmt.Println("getDbEnv err")
	}
	db, err := sql.Open("postgres", "host="+df.Host+" user="+df.User+" password="+df.Password+" dbname="+df.Database+" port="+df.Port+" sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic(err.Error())
	}
	return db
}
