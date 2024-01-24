package services

import (
	"errors"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
)

type AccountServices struct{}

func (svc AccountServices) GetAccount(userID uint) (account models.Accounts, err error) {
	// Get account
	result := initializers.DB.Where("user_id=?", userID).First(&account)
	if result.Error != nil {
		err := errors.New("failed to get account")
		return account, err
	}

	// Return account
	return account, nil

}
