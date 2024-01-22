package controllers

import (
	"net/http"

	"github.com/Eco-Led/EcoLed-Back_test/services"
	"github.com/gin-gonic/gin"
)

type UserControllers struct{}

var userService = new(services.UserServices)

func (ctr UserControllers) Login(c *gin.Context) {
	var loginForm services.LoginForm

	// Bind JSON
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Login(service)
	user, token, err := userService.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response user, token
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})


}

func (ctr UserControllers) Register(c *gin.Context) {
	var registerForm services.RegisterForm

	// Bind JSON
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Register(service)
	err := userService.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Register Success",
	})


}

func (ctr UserControllers) Logout(c *gin.Context) {
}
