package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string

	// usersprofiles: OneToOne mapping
	Profiles Profiles `gorm:"foreignkey:User_id"`
	// accounts: OneToOne mapping
	Accounts Accounts `gorm:"foreignkey:User_id"`
}
