package jwt

import (
	"time"

	"github.com/designsbysm/server-go/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func Encode(id uuid.UUID, role database.Role) (string, error) {
	defaulDuration := 8
	if role.IsAdmin {
		defaulDuration = 24 * 365
	}
	expiration := time.Now().Add(time.Hour * time.Duration(defaulDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role.Name,
		"exp":  expiration,
	})

	secretKey := []byte(viper.GetString("jwt.secret"))

	return token.SignedString(secretKey)
}
