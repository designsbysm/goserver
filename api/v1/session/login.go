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
		Email: request.Email,
	}
	err = user.Read(database.PreloadRole)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	err = user.ValidatePassword(request.Password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	session := database.Session{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role.Name,
	}
	if err = session.Upsert(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err := jwt.Encode(session.ID, *user.Role)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Token = token
	if err := session.Upsert(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, session)
}
