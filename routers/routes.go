package routers

import (
	"github.com/Eco-Led/EcoLed-Back_test/controllers/auth"

	"github.com/gin-gonic/gin"
)

func GoogleAuthRoutes(router *gin.Engine, apiVersion string) {
	router.GET(apiVersion+"auth/google/login", auth.GoogleLoginHandler)
	router.GET(apiVersion+"auth/google/callback", auth.GoogleAuthCallback)
}

func AuthRoutes(router *gin.Engine, apiVersion string) {
	router.POST(apiVersion + "/login" /* Not Implemented */)
	router.POST(apiVersion + "/register" /* Not Implemented */)
	router.GET(apiVersion + "/logout" /* Not Implemented */)
}

func RouterSetupV1() *gin.Engine {
	r := gin.Default()

	apiVersion := "/api/v1"
	r.Group(apiVersion)
	{
		AuthRoutes(r, apiVersion)
		GoogleAuthRoutes(r, apiVersion)
	}

	return r
}
