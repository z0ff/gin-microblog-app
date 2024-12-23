package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex, not null, type:varchar(100)"`
	Posts []Post `gorm:"foreignKey:UserID"`
}
