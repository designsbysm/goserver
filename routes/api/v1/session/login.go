package session

import (
	"net/http"

	"github.com/designsbysm/goserver/database"
	role "github.com/designsbysm/goserver/routes/api/v1/role/orm"
	"github.com/designsbysm/goserver/routes/api/v1/session/orm"
	"github.com/designsbysm/goserver/utilities"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var query database.User

	err := c.BindJSON(&query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if query.Email == "" || query.Password == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	password := query.Password
	query.Password = ""

	user, err := orm.ReadUser(query)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	err = user.ValidatePassword(password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	role, err := role.ReadRole(int(user.RoleID))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session := database.Session{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      role.Name,
	}

	token, err := utilities.JWTEncode(session)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Token = token

	err = orm.UpdateUser(user, session.Token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, session)
}
