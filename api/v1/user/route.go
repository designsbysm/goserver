package user

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/user")
	{
		// group.Use(middleware.AuthorizeJWT())

		group.GET("/", list)    // TODO: add admin check
		group.POST("/", create) // TODO: add admin check

		// id := r.Group("/:id")
		// {
		group.GET("/:id", read)      // TODO: admin or user
		group.DELETE("/:id", delete) // TODO: add admin check
		group.PUT("/:id", update)    // TODO: admin or user
		// }
	}
}
