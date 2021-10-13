package jwt

import (
	"time"

	"github.com/designsbysm/server-go/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func Encode(session database.Session) (string, error) {
	secretKey := []byte(viper.GetString("server.jwt.secret"))

	defaulDuration := 8
	if session.Role == "admin" {
		defaulDuration = 24 * 365
	}

	expiration := time.Now().Add(time.Hour * time.Duration(defaulDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   session.ID,
		"role": session.Role,
		"exp":  expiration,
	})

	return token.SignedString(secretKey)
}
