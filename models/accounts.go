package models

import (
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Bank            string
	Number          int64 `gorm:"uniqueIndex"`
	Balance         int64
	Total_eco_score float64

	// users: OneToOne mapping
	User_id uint `gorm:"uniqueIndex"`

	// paylogs: OneToMany mapping
	Paylogs []Paylogs `gorm:"foreignkey:Account_id"`
}
