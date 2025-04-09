package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex not null, type:varchar(32)"`
	Email string `gorm:"uniqueIndex, not null, type:varchar(100)"`
}
