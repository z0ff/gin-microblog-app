package model

import (
	"gorm.io/gorm"
)

type UserAuth struct {
	gorm.Model
	UserID   uint   `gorm:"uniqueIndex;not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Password string `gorm:"not null, type:varchar(255)"`
}
