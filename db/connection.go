package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/db/seeder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
	err        error
)

type Database struct {
	Connection *gorm.DB
}

// DB接続の初期化
func NewDatabase() (*Database, error) {
	// 環境変数の読み込み
	err := godotenv.Load(".env")

	// DB接続
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db_conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db_conn.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		return nil, err
	}

	seeder.UserSeeder(db_conn)

	connection = db_conn

	return &Database{Connection: db_conn}, nil

}

// DB接続の取得
func GetConnection() *gorm.DB {
	return connection
}
