package database

import (
	"github.com/designsbysm/logger/v2"
	"github.com/spf13/viper"
)

func populateDB() {
	setting := Setting{
		Key: "database",
	}
	if err := setting.Read(); err != nil {
		logger.Error(err)
	} else if setting.GetBool("populated") {
		return
	}

	role := Role{
		Name:    "admin",
		IsAdmin: true,
	}
	if err := role.Create(); err != nil {
		logger.Error(err)
	}

	user := User{
		FirstName: "Scott",
		LastName:  "Matthews",
		Email:     "scott@designsbysm.com",
		Password:  viper.GetString("db.user.password"),
		RoleID:    role.ID,
	}
	if err := user.Create(); err != nil {
		logger.Error(err)
	}

	role = Role{
		Name:    "user",
		IsAdmin: false,
	}
	if err := role.Create(); err != nil {
		logger.Error(err)
	}

	user = User{
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob@designsbysm.com",
		Password:  viper.GetString("db.user.password"),
		RoleID:    role.ID,
	}
	if err := user.Create(); err != nil {
		logger.Error(err)
	}

	setting.Value["populated"] = true
	if err := setting.Upsert(); err != nil {
		logger.Error(err)
	}
}
