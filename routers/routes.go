package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Eco-Led/EcoLed-Back_test/controllers"
	"github.com/Eco-Led/EcoLed-Back_test/middlewares"
)

var userController = new(controllers.UserControllers)
var profileController = new(controllers.ProfileControllers)
var imageController = new(controllers.ImageControllers)
var accountController = new(controllers.AccountControllers)
var paylogController = new(controllers.PaylogControllers)

func AuthRoutes(router *gin.Engine, apiVersion string) {
	router.POST(apiVersion+"/login", userController.Login)
	router.POST(apiVersion+"/register", userController.Register)
	router.GET(apiVersion+"/logout", userController.Logout)
}

func ProfileRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.PUT(apiVersion+"/profile", profileController.UpdateProfile)
	router.GET(apiVersion+"/profile", profileController.GetProfile)
}

func ImageRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.POST(apiVersion+"/upload", imageController.UploadImage)
}

func AccountRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.GET(apiVersion+"/account", accountController.GetAccount)
}

func PaylogRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.POST(apiVersion+"/paylog", paylogController.CreatePaylog)
	router.PUT(apiVersion+"/paylog/:paylogID", paylogController.UpdatePaylog)
}

func RouterSetupV1() *gin.Engine {
	r := gin.Default()

	apiVersion := "/api/v1"
	r.Group(apiVersion)
	{
		AuthRoutes(r, apiVersion)
		ProfileRoutes(r, apiVersion)
		ImageRoutes(r, apiVersion)
		AccountRoutes(r, apiVersion)
		PaylogRoutes(r, apiVersion)
	}

	return r
}
