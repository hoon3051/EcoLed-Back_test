package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title string
	Body  string
	Image string

	// From Users: OneToMany mapping
	User_id uint `gorm:"foreignkey:User_id"`
}
