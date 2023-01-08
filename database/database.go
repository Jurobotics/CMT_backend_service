package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type DefaultModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

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

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conntection Opend to Database")

}