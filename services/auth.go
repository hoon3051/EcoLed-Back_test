package services

import (
	"errors"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email      string
	Nickname   string
	created_at string
}
type LoginForm struct {
	Email    string `form: "email" json: "email" binding: "required, email"`
	Password string `form: "password" json: "password" binding: "required, min=6, max=30"`
}

type UserServices struct{}

var userModel = new(models.Users)
var profileModel = new(models.Profiles)
var tokenService = new(TokenServices)

func (svc UserServices) Login(loginForm LoginForm) (user User, token Token, err error) {
	initializers.DB.First(&userModel, "email=?", loginForm.Email)
	initializers.DB.First(&profileModel, "user_id=?", userModel.ID)

	if userModel.ID == 0 || profileModel.ID == 0 {
		err := errors.New("Data does not exist in DB")
		return user, token, err
	}

	user = User{
		Email:      userModel.Email,
		Nickname:   profileModel.Nickname,
		created_at: userModel.CreatedAt.String(),
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginForm.Password))
	if err != nil {
		err := errors.New("Invalid Password")
		return user, token, err
	}

	td, err := tokenService.CreateToken(int64(userModel.ID))
	if err != nil {
		return user, token, err
	}

	err = tokenService.SaveToken(int64(userModel.ID), td)
	if err != nil {
		return user, token, err
	}

	token = Token{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}

	return user, token, err
}

// func (svc UserServices) Register(loginForm RegisterForm) {

// }

func (svc UserServices) Logout() {

}
