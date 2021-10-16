package user

import (
	"net/http"

	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/server-go/middleware"
	"github.com/designsbysm/server-go/tools"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/user")
	{
		group.Use(middleware.AuthorizeJWT())
		group.GET("/", middleware.AuthorizeAdmin(), list)
		group.POST("/", middleware.AuthorizeAdmin(), create)

		id := group.Group("/:id")
		{
			id.GET("/", adminOrUser(), read)
			id.PUT("/", adminOrUser(), update)
			id.DELETE("/", middleware.AuthorizeAdmin(), delete)
		}
	}
}

func adminOrUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		user := data.(database.User)

		if err := user.Read(database.PreloadRole); err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		} else if user.Role.IsAdmin {
			return
		}

		id, err := tools.GetIDParam(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if user.ID != id {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
