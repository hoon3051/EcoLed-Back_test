package services

import (
	"errors"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
)

// Profile service's input value (from body)
type ProfileForm struct {
	Nickname 		string `form: "nickname" json: "nickname" binding: "required, min=2, max=30"`
	Age	  	 		int    `form: "age" json: "age" binding: "required, min=0, max=150"`
	Introduce 		string `form: "introduce" json: "introduce" binding: "required"`
}

type ProfileServices struct{}

func (svc ProfileServices) UpdateProfile(userID uint, profileForm ProfileForm) (err error) {
	// Update profile
	result := initializers.DB.Model(&models.Profiles{}).Where("user_id=?", userID).Updates(models.Profiles{
		Nickname: profileForm.Nickname,
		Age: profileForm.Age,
		Introduce: profileForm.Introduce,
	})
	if result.Error != nil {
		err := errors.New("Failed to update profile")
		return err
	}

	return nil

}

func (svc ProfileServices) GetProfile(userID uint) (profile models.Profiles, err error) {
	// Get profile
	result := initializers.DB.Where("user_id=?", userID).First(&profile)
	if result.Error != nil {
		err := errors.New("Failed to get profile")
		return profile, err
	}

	return profile, nil
	
}