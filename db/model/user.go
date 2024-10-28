package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex, not null, type:varchar(100)"`
}
