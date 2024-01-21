package initializers

import (
	"os"

	"github.com/Eco-Led/EcoLed-Back_test/logger"
	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Println("Failed to load .env file")
		os.Exit(-1)
	}
}
