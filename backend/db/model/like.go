package model

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID uint
	PostID uint
	User   User `gorm:"foreignKey:UserID"`
	Post   Post `gorm:"foreignKey:PostID"`
}
