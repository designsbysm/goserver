package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func gormLogger(level int) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", 0),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.LogLevel(level),
			SlowThreshold:             200 * time.Millisecond,
		},
	)
}
