package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type request struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	RoleID    uint
}

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

	request := request{}
	err = c.BindJSON(&request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}

	if request.LastName != "" {
		user.LastName = request.LastName
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.RawPassword = request.Password
	}

	if request.RoleID != 0 {
		user.RoleID = request.RoleID
		user.Role = nil
	}

	err = user.Update()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
