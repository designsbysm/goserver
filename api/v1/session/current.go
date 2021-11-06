package session

import (
	"errors"
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

func current(c *gin.Context) {
	data, ok := c.Get("user")
	if !ok {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user data"))
		return
	}
	user := data.(database.User)

	session := database.Session{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role.Name,
	}

	c.JSON(http.StatusOK, session)
}
