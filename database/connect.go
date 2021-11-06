package database

import (
	"github.com/designsbysm/server-go/database/incident"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	if DB, err = gorm.Open(
		postgres.Open(
			viper.GetString("db.connection"),
		),
		&gorm.Config{
			Logger: gormLogger(viper.GetInt("gorm.level")),
		}); err != nil {
		return
	}

	if err := DB.AutoMigrate(
		&Role{},
		&Setting{},
		&User{},
		&Session{},
	); err != nil {
		return err
	}
	populateDB()

	if err := DB.AutoMigrate(
		&incident.InIncident{},
		&incident.InMember{},
		&incident.InMemberLevel{},
		&incident.InPatient{},
		&incident.InPCR{},
		&incident.InUnit{},
	); err != nil {
		return err
	}

	return
}
