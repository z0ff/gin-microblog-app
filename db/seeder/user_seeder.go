package seeder

import (
	"github.com/z0ff/microblog-backend/db/model"
	"gorm.io/gorm"
)

// UserSeeder はUserテーブルの初期データを挿入する
func UserSeeder(db *gorm.DB) {
	users := []model.User{
		{
			Name:        "john123",
			DisplayName: "John",
			Email:       "john@example.com",
			Password:    "password",
		},
	}

	db.Exec("DELETE from users")
	db.Exec("ALTER TABLE users auto_increment = 1")

	for _, user := range users {
		db.Create(&user)
	}
}
