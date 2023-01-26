package database

import (
	"fmt"
	"juro-go/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func SetupDatabase(path string) {
	var err error
	var config gorm.Config

	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		config = gorm.Config{}
	} else {
		config = gorm.Config{
			Logger: logger.Default.LogMode((logger.Silent)),
		}
	}

	DB, err = gorm.Open(sqlite.Open(path), &config)

	DB.AutoMigrate(&models.Recipe{}, &models.Ingredient{}, &models.Servo{}, &models.Ingredients{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conntection Opend to Database")

}
