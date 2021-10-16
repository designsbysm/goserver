package user

import (
	"net/http"

	"github.com/designsbysm/logger/v2"
	"github.com/designsbysm/server-go/database"
	"github.com/gin-gonic/gin"
)

// type createRequest struct {
// 	Password string
// }

func create(c *gin.Context) {

	// id, err := tools.GetIDParam(c)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// user := database.User{
	// 	ID: id,
	// }

	// err = user.Read(database.PreloadRole)
	// if err != nil {
	// 	status := http.StatusInternalServerError
	// 	if err == gorm.ErrRecordNotFound {
	// 		status = http.StatusBadRequest
	// 	}

	// 	c.AbortWithError(status, err)
	// 	return
	// }

	// user := database.User{}
	// err := c.BindJSON(&user)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// request := request{}
	user := database.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// user := database.User{}

	// if request.FirstName != "" {
	// 	user.FirstName = request.FirstName
	// }

	// if request.LastName != "" {
	// 	user.LastName = request.LastName
	// }

	// if request.Email != "" {
	// 	user.Email = request.Email
	// }

	// if request.Password != "" {
	// 	user.RawPassword = request.Password
	// }

	// if request.RoleID != 0 {
	// 	user.RoleID = request.RoleID
	// 	user.Role = nil
	// }

	// logger.Struct(user)

	// user.RawPassword = request.Password
	// b, err := ioutil.ReadAll(c.Request.Body)
	// logger.Debug(c.Request.)

	logger.Struct(user)

	err = user.Create()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
