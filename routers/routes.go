package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Eco-Led/EcoLed-Back_test/controllers"
)

var userController = new(controllers.UserControllers)

func AuthRoutes(router *gin.Engine, apiVersion string) {
	router.POST(apiVersion+"/login", userController.Login)
	router.POST(apiVersion+"/register", userController.Register)
	router.GET(apiVersion+"/logout", userController.Logout)
}

func RouterSetupV1() *gin.Engine {
	r := gin.Default()

	apiVersion := "/api/v1"
	r.Group(apiVersion)
	{
		AuthRoutes(r, apiVersion)
	}

	return r
}
