package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	user := database.User{}

	list, err := user.List(0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, list)
}
