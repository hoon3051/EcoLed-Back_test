package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Eco-Led/EcoLed-Back_test/controllers"
	"github.com/Eco-Led/EcoLed-Back_test/middlewares"
)

var userController = new(controllers.UserControllers)
var profileController = new(controllers.ProfileControllers)
var profileimageController = new(controllers.ProfileImageControllers)
var accountController = new(controllers.AccountControllers)
var paylogController = new(controllers.PaylogControllers)
var rankingController = new(controllers.RankingControllers)
var postController = new(controllers.PostControllers)

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

func ProfileImageRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.POST(apiVersion+"/profileimage", profileimageController.UploadProfileImage)
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

func RankingRoutes(router *gin.Engine, apiVersion string) {
	router.GET(apiVersion+"/ranking", rankingController.GetRanking)
}

func PostRoutes(router *gin.Engine, apiVersion string) {
	router.Use(middlewares.AuthToken())
	router.POST(apiVersion+"/post", postController.CreatePost)
	router.GET(apiVersion+"/post", postController.GetUserPosts)
	router.GET(apiVersion+"/post/:postID", postController.GetOnePost)
	router.PUT(apiVersion+"/post/:postID", postController.UpdatePost)
}

func RouterSetupV1() *gin.Engine {
	r := gin.Default()

	apiVersion := "/api/v1"
	r.Group(apiVersion)
	{
		AuthRoutes(r, apiVersion)
		ProfileRoutes(r, apiVersion)
		ProfileImageRoutes(r, apiVersion)
		AccountRoutes(r, apiVersion)
		PaylogRoutes(r, apiVersion)
		RankingRoutes(r, apiVersion)
		PostRoutes(r, apiVersion)
	}

	return r
}
