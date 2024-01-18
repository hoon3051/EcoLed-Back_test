package models

import(
	"gorm.io/gorm"
)

type Paylog_details struct{
	gorm.Model
	Name 		string
	Place 		string
	Material 	string
	Eco_score 		float64

	// pay_logs로부터 OneToOne mapping
	Paylog_id uint `gorm:"uniqueIndex"`
}