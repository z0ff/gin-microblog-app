package seeder

import (
	"github.com/z0ff/hello-go/db/model"
	"gorm.io/gorm"
)

// UserSeeder はUserテーブルの初期データを挿入する
func UserSeeder(db *gorm.DB) {
	users := []model.User{
		{
			Name:  "John",
			Email: "john@example.com",
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
