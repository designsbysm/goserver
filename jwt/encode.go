package jwt

import (
	"time"

	"github.com/designsbysm/server-go/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func Encode(user *database.User) (Session, error) {
	session := Session{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role.Name,
	}
	secretKey := []byte(viper.GetString("server.jwt.secret"))

	defaulDuration := 8
	if user.Role.IsAdmin {
		defaulDuration = 24 * 365
	}

	expiration := time.Now().Add(time.Hour * time.Duration(defaulDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   session.ID,
		"role": session.Role,
		"exp":  expiration,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return session, err
	}

	user.AuthToken = tokenString
	session.Token = tokenString

	return session, nil
}
