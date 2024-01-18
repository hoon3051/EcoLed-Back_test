package models

import(
	"gorm.io/gorm"
)

type Pay_logs struct{
	gorm.Model
	Date string
	Time string
	Content string
	Cost int64

	// accounts로부터 OneToMany mapping	
	Account_id uint `gorm:"foreignkey:Account_id"`

	// eco_stamps로 OneToOne mapping
	Paylog_details Paylog_details `gorm:"foreignkey:Paylog_id"`
	
}