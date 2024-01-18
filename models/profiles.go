package models

import (
	"gorm.io/gorm"
)

type Profiles struct {
	gorm.Model
	Nickname      string `gorm:"uniqueIndex"`
	Age           int16
	Profile_image string
	Introduce     string

	// users: OneToOne mapping
	User_id uint `gorm:"uniqueIndex"`
}
