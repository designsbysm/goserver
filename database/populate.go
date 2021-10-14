package database

import (
	"github.com/designsbysm/logger/v2"
	"github.com/spf13/viper"
)

func populateDB() {
	roleAdmin := Role{
		Name:    "admin",
		IsAdmin: true,
	}
	createRecord(&roleAdmin, roleAdmin)

	userAdmin := User{
		FirstName: "Scott",
		LastName:  "Matthews",
		Email:        "scott@designsbysm.com",
		Password:  viper.GetString("db.user.password"),
		RoleID:    roleAdmin.ID,
	}
	createRecord(&userAdmin, userAdmin)

	roleUser := Role{
		Name:    "user",
		IsAdmin: true,
	}
	createRecord(&roleUser, roleUser)

	userUser := User{
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob@designsbysm.com",
		Password:  viper.GetString("db.user.password"),
		RoleID:    roleUser.ID,
	}
	createRecord(&userUser, userUser)
}

func createRecord(data interface{}, query interface{}) {
	err := DB.FirstOrCreate(data, query).Error
	if err != nil {
		logger.Error(err)
	}
}
