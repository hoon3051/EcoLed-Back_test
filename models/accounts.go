package models

import (
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Name            string
	Balance         int64
	Total_ecoscore  float64

	// From Users: OneToOne mapping
	User_id uint `gorm:"uniqueIndex"`

	// To Paylogs: OneToMany mapping
	Paylogs []Paylogs `gorm:"foreignkey:Account_id"`
}
