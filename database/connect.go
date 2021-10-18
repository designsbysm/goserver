package database

import (
	"github.com/designsbysm/server-go/database/incident"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	DB, err = gorm.Open(
		postgres.Open(
			viper.GetString("db.connection"),
		),
		&gorm.Config{
			Logger: gormLogger(viper.GetInt("gorm.level")),
		})
	if err != nil {
		return
	}

	DB.AutoMigrate(
		&Role{},
		&Setting{},
		&User{},
		&Session{},
	)
	populateDB()

	DB.AutoMigrate(
		&incident.Incident{},
		&incident.Member{},
		&incident.MemberLevel{},
		&incident.Patient{},
		&incident.PCR{},
		&incident.Unit{},
	)

	return
}
