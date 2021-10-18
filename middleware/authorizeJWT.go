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

		session := database.Session{
			UserID: uint(claims["id"].(float64)),
			Token:  token,
		}
		if err := session.Read(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user := database.User{
			ID: session.UserID,
		}
		if err = user.Read(database.PreloadRole); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
