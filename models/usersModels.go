package models

import(
	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	Email string `gorm:"uniqueIndex"`
	Password string

	// users_profiles로 OneToOne mapping
	Users_profiles User_profiles `gorm:"foreignkey:User_id"`
	// accounts로 OneToOne mapping
	Accounts Accounts `gorm:"foreignkey:User_id"`
}