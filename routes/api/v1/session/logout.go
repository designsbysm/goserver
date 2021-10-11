package session

import (
	"errors"
	"net/http"

	"github.com/designsbysm/goserver/database"
	"github.com/designsbysm/goserver/routes/api/v1/session/orm"
	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	data, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user data"))
		return
	}
	user := data.(database.User)

	err := orm.UpdateUser(user, "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
