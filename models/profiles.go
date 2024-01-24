package models

import (
	"gorm.io/gorm"
)

// Age, Profile_image, Introduce can be null so use pointer(*) type
type Profiles struct {
	gorm.Model
	Nickname      string `gorm:"unique"`
	Age           int
	Profile_image string
	Introduce     string

	// From Users: OneToOne mapping
	User_id uint `gorm:"uniqueIndex"`
}
