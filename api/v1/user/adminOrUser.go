package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func adminOrUser() gin.HandlerFunc {
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
		} else if user.Role.IsAdmin {
			return
		}

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if user.ID != id {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
