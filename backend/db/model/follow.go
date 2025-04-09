package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserID      uint
	FollowingID uint
	User        User `gorm:"foreignKey:UserID"`
	Following   User `gorm:"foreignKey:FollowingID"`
}
