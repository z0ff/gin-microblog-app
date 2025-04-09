package model

import (
	"gorm.io/gorm"
)

type AdminAuth struct {
	gorm.Model
	AdminID  uint   `gorm:"uniqueIndex;not null"`
	Admin    Admin  `gorm:"foreignKey:AdminID"`
	Password string `gorm:"not null, type:varchar(255)"`
}
