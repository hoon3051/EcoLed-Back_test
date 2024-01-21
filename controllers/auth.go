package controllers

import (
	"github.com/Eco-Led/EcoLed-Back_test/services"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `form: "email" json: "email" binding: "required, email"`
	Password string `form: "password" json: "password" binding: "required, min=6, max=30"`
}

type UserControllers struct{}

var userService = new(services.UserServices)

func (ctr UserControllers) Login() (c *gin.Context) {

}

func (ctr UserControllers) Register() (c *gin.Context) {

}

func (ctr UserControllers) Logout() (c *gin.Context) {

}
