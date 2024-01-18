package controllers

import (
	"github.com/Eco-Led/EcoLed-Back_test/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/google/login", auth.GoogleLoginHandler)
	router.GET("/auth/google/callback", auth.GoogleAuthCallback)
}
