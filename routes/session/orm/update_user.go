package orm

import (
	"github.com/designsbysm/goserver/database"
)

func UpdateUser(user database.User, token string) error {
	return database.DB.Model(&user).Update("auth_token", token).Error
}
