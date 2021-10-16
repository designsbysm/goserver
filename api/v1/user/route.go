package user

import (
	"github.com/designsbysm/server-go/middleware"
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
