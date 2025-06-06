package seeder

import (
	"github.com/z0ff/microblog-backend/db/model"
	"gorm.io/gorm"
	"github.com/z0ff/microblog-backend/utils/crypto"
)

// UserSeeder はUserテーブルの初期データを挿入する
func UserSeeder(db *gorm.DB) {
	pw_hash, _ := crypto.HashPassword("password")
	users := []model.User{
		{
			Name:        "john123",
			DisplayName: "John",
			Email:       "john@example.com",
		},
	}
	user_auths := []model.UserAuth{
		{
			UserID:   1,
			Password: pw_hash,
		},
	}

	db.Exec("DELETE from user_auths")
	db.Exec("DELETE from users")
	db.Exec("ALTER TABLE users auto_increment = 1")

	for _, user := range users {
		db.Create(&user)
	}
	for _, user_auth := range user_auths {
		db.Create(&user_auth)
	}
}
