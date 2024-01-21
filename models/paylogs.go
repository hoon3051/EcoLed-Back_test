package models

import (
	"gorm.io/gorm"
)

type Paylogs struct {
	gorm.Model
	Date    	string
	Time    	string
	Content 	string
	Cost    	int64
	Name		string
	Place		string
	Material	string
	Ecoscore 	float64

	// From accounts: OneToMany mapping
	Account_id uint `gorm:"foreignkey:Account_id"`
}
