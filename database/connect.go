package database

import (
	"fmt"

	"github.com/designsbysm/server-go/database/incident"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	connection := viper.GetString("URL_DATABASE")
	if connection == "" {
		user := viper.GetString("POSTGRES_USER")
		password := viper.GetString("POSTGRES_PASSWORD")
		port := viper.GetString("PORT_POSTGRES")
		db := viper.GetString("POSTGRES_DB")

		connection = fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, port, db)
	}

	if DB, err = gorm.Open(
		postgres.Open(connection),
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
