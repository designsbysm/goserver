package middleware

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		user := data.(database.User)

		if err := user.Read(database.PreloadRole); err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		} else if !user.Role.IsAdmin {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
