package middleware

import (
	"net/http"
	"strings"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		id, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		session := database.Session{
			ID:    id,
			Token: token,
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
