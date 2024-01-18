package models

import (
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Bank 	  	string
	Number      int64 `gorm:"uniqueIndex"`
	Balance     int64
	Total_eco_score float64

	// users로부터 OneToOne mapping
	User_id uint `gorm:"uniqueIndex"`

	// pay_logs로 OneToMany mapping
	Pay_logs []Pay_logs `gorm:"foreignkey:Account_id"`
}
