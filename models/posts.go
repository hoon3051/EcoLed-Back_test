package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title      	  string `gorm:"uniqueIndex"`
	Body          string
	Image 	   	  string

	// From Users: OneToMany mapping
	User_id uint `gorm:"foreignkey:User_id"`
}