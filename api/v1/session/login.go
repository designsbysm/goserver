package session

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/jwt"

	"github.com/gin-gonic/gin"
)

type request struct {
	Email    string
	Password string
}

func login(c *gin.Context) {
	request := request{}
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if request.Email == "" || request.Password == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	user := database.User{
		Email:       request.Email,
		RawPassword: request.Password,
	}
	err = user.Read(database.PreloadRole)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	err = user.ValidatePassword()
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	session, err := jwt.Encode(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = user.Update()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, session)
}
