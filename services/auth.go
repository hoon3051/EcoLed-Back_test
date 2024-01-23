package services

import (
	"errors"
	"regexp"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
	"golang.org/x/crypto/bcrypt"
)

// Login service's return value
type User struct {
	Email      string
	Nickname   string
	created_at string
}

// Login service's input value (from body)
type LoginForm struct {
	Email    string `form: "email" json: "email" binding: "required, email"`
	Password string `form: "password" json: "password" binding: "required, min=6, max=30"`
}

// Register service's input value (from body)
type RegisterForm struct {
	Email    	string `form: "email" json: "email" binding: "required, email"`
	Password 	string `form: "password" json: "password" binding: "required, min=6, max=30"`
	Nickname 	string `form: "nickname" json: "nickname" binding: "required, min=2, max=30"`
	Accountname string `form: "accountname" json: "accountname" binding: "required, min=2, max=30"`
}

type UserServices struct{}



func (svc UserServices) Login(loginForm LoginForm) (user User, token Token, err error) {
	//call by value (not call by reference)
	var userModel = models.Users{}
	var profileModel = models.Profiles{}

	//call by reference
	var tokenService = new(TokenServices)

	//From controller, binding value is received. So, check whether the value is valid.
	initializers.DB.First(&userModel, "email=?", loginForm.Email)
	initializers.DB.First(&profileModel, "user_id=?", userModel.ID)

	// If the value is not valid, return error
	if userModel.ID == 0 || profileModel.ID == 0 {
		err := errors.New("Data does not exist in DB")
		return user, token, err
	}

	// Set return value (user)
	user = User{
		Email:      userModel.Email,
		Nickname:   profileModel.Nickname,
		created_at: userModel.CreatedAt.String(),
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginForm.Password))
	if err != nil {
		err := errors.New("Invalid Password")
		return user, token, err
	}

	// Create token
	td, err := tokenService.CreateToken(int64(userModel.ID))
	if err != nil {
		return user, token, err
	}

	// Save token
	err = tokenService.SaveToken(int64(userModel.ID), td)
	if err != nil {
		return user, token, err
	}

	// Set return value (token)
	token = Token{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}

	return user, token, err
}

func (svc UserServices) Register(registerForm RegisterForm) (err error) {
	//call by value (not call by reference)
	var userModel = models.Users{}
	var profileModel = models.Profiles{}

	// Check whether the email is unique
	initializers.DB.First(&userModel, "email=?", registerForm.Email)
	if userModel.ID != 0 {
		err := errors.New("Email already exists")
		return err
	}

	//Check whether the nickname is unique
	initializers.DB.First(&profileModel, "nickname=?", registerForm.Nickname)
	if profileModel.ID != 0 {
		err := errors.New("Nickname already exists")
		return err
	}

	// Check whether the email is valid
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    if regex.MatchString(registerForm.Email) == false {
		err := errors.New("Invalid email")
		return err
	}

	// Check whether the nickname is valid
	if len(registerForm.Nickname) < 2 || len(registerForm.Nickname) > 30 {
		err := errors.New("Nickname must be between 2 and 30 characters")
		return err
	}

	// Check whether the password is valid
	if len(registerForm.Password) < 4 || len(registerForm.Password) > 30 {
		err := errors.New("Password must be between 4 and 30 characters")
		return err
	}

	

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), bcrypt.DefaultCost)
	if err != nil {
		err := errors.New("Failed to hash password")
		return err
	}

	// Create user
	user := models.Users{
		Email:    registerForm.Email,
		Password: string(hashedPassword),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		err := errors.New("Failed to create user")
		return err
	}

	// Create profile
	result = initializers.DB.Create(&models.Profiles{
		Nickname: registerForm.Nickname,
		User_id:   user.ID,
	})
	if result.Error != nil {
		err := errors.New("Failed to create profile")
		return err
	}

	// Create account
	result = initializers.DB.Create(&models.Accounts{
		Name: registerForm.Accountname,
		User_id:   user.ID,
	})
	if result.Error != nil {
		err := errors.New("Failed to create account")
		return err
	}

	return nil
}

func (svc UserServices) Logout() {

}
