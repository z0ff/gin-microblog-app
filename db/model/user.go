package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex not null, type:varchar(32)"`
	DisplayName string `gorm:"not null, type:varchar(100)"`
	Password    string `gorm:"not null, type:varchar(255)"`
	Email       string `gorm:"uniqueIndex, not null, type:varchar(100)"`
	Posts       []Post `gorm:"foreignKey:UserID"`
}
