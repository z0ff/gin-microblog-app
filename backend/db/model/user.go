package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex not null, type:varchar(32)"`
	DisplayName string `gorm:"not null, type:varchar(100)"`
	Email       string `gorm:"uniqueIndex, not null, type:varchar(100)"`
	Posts       []Post `gorm:"foreignKey:UserID"`
	SuspendedAt *time.Time
}
