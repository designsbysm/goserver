package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func update(c *gin.Context) {
	id, err := tools.GetIDParam(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := database.User{
		ID: id,
	}
	err = user.Read(database.PreloadRole)
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusBadRequest
		}

		c.AbortWithError(status, err)
		return
	}
	currentRole := user.RoleID

	err = c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if user.RawPassword != "" {
		user.Password = ""
	}

	if user.RoleID != currentRole {
		role := database.Role{
			ID: user.RoleID,
		}
		if err := role.Read(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user.Role = &role
	}

	err = user.Update()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
