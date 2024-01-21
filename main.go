package main

import (
	"os"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/logger"
	"github.com/Eco-Led/EcoLed-Back_test/routers"
)

func init() {
	initializers.LoadDotEnv()
	initializers.InitDB()
}

func main() {
	router := routers.RouterSetupV1()


	port := os.Getenv("PORT")
	logger.Info.Println("Server listening in port: ", port)
	router.Run(":" + port) // TODO : add SSL, TLS connect
}
