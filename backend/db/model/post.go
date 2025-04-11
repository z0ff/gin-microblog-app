package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	User    User   `gorm:"foreignKey:UserID"`
	Content string `gorm:"not null, type:varchar(255)"`
}
