package initializers

import (
	"fmt"
	"os"

	"github.com/Eco-Led/EcoLed-Back_test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to DB")
	}

}

func SyncDB() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.User_profiles{})
	DB.AutoMigrate(&models.Accounts{})
	DB.AutoMigrate(&models.Pay_logs{})
	DB.AutoMigrate(&models.Paylog_details{})
}
