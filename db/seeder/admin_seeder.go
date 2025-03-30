package seeder

import (
	"github.com/z0ff/microblog-backend/db/model"
	"github.com/z0ff/microblog-backend/utils/crypto"
	"gorm.io/gorm"
)

// AdminSeeder はAdminテーブルの初期データを挿入する
func AdminSeeder(db *gorm.DB) {
	pw_hash, _ := crypto.HashPassword("password")
	admins := []model.Admin{
		{
			Name:  "john123",
			Email: "john@example.com",
		},
	}
	admin_auths := []model.AdminAuth{
		{
			AdminID:  1,
			Password: pw_hash,
		},
	}

	db.Exec("DELETE from admin_auths")
	db.Exec("DELETE from admins")
	db.Exec("ALTER TABLE users auto_increment = 1")

	for _, admin := range admins {
		db.Create(&admin)
	}
	for _, admin_auth := range admin_auths {
		db.Create(&admin_auth)
	}
}
