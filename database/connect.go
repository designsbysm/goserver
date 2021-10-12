package database

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.LogLevel(viper.GetInt("gorm.level")),
			SlowThreshold:             time.Second,
		},
	)

	var err error
	DB, err = gorm.Open(
		postgres.Open(
			viper.GetString("db.connection"),
		),
		&gorm.Config{
			Logger: customLogger,
		})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&Role{}, &User{})
	populateDB()

	return err
}
