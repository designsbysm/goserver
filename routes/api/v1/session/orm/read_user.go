package orm

import (
	"github.com/designsbysm/server-go/database"
)

func ReadUser(query database.User) (database.User, error) {
	var user database.User

	err := database.DB.First(&user, query).Error
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
