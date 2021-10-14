package database

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

func gormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", 0),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.LogLevel(viper.GetInt("gorm.level")),
			SlowThreshold:             200 * time.Millisecond,
		},
	)
}
