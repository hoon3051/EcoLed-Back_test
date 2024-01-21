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

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, token, err := userService.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})

	return
}

func (ctr UserControllers) Register(c *gin.Context) {
	return
}

func (ctr UserControllers) Logout(c *gin.Context) {
	return
}
