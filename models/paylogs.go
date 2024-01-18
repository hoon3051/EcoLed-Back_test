package models

import (
	"gorm.io/gorm"
)

type Paylogs struct {
	gorm.Model
	Date    string
	Time    string
	Content string
	Cost    int64

	// accounts: OneToMany mapping
	Account_id uint `gorm:"foreignkey:Account_id"`
}
