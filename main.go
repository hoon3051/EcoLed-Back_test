package main

import (
	"fmt"
	"os"

	"github.com/Eco-Led/EcoLed-Back_test/controllers"
	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func sendEmail() {
	smtp.PlainAuth(
		"",
		"hoon30512329@gmail.com",
		""
	)

func main() {
	fmt.Println("EcoLed!!")

	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credPath == "" {
		fmt.Println("GOOGLE_APPLICATION_CREDENTIALS is not set")
	} else {
		fmt.Println("GOOGLE_APPLICATION_CREDENTIALS:", credPath)
	}

	router := gin.Default()

	controllers.AuthRoutes(router)

	router.Run()
}
