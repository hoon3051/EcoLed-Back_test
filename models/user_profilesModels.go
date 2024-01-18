package models

import(
	"gorm.io/gorm"
)

type User_profiles struct{
	gorm.Model
	Nickname string `gorm:"uniqueIndex"`
	Age int16
	Profile_image string 
	Introduce string

	// users로부터 OneToOne mapping	
	User_id uint `gorm:"uniqueIndex"`
}