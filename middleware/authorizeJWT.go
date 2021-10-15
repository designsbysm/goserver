package middleware

import (
	"net/http"
	"strings"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/jwt"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := jwt.Decode(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user := database.User{
			AuthToken: token,
		}
		user.ID = uint(claims["id"].(float64))
		if err = user.Read(database.PreloadRole); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
