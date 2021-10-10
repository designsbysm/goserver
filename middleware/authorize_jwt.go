package middleware

import (
	"net/http"
	"strings"

	"github.com/designsbysm/goserver/database"
	"github.com/designsbysm/goserver/utilities"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := utilities.JWTDecode(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user database.User
		var query database.User

		query.ID = uint(claims["id"].(float64))
		query.AuthToken = token

		err = database.DB.Joins("Role").First(&user, query).Error
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
