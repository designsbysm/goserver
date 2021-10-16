package user

import (
	"net/http"

	"github.com/designsbysm/logger/v2"
	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/tools"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func read(c *gin.Context) {
	// ui := c.GetUint("id")
	// logger.Debug(ui)

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

	u := structs.Map(user)
	logger.Struct(u)

	c.JSON(http.StatusOK, user)
}
