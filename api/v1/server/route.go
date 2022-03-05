package server

import (
	"github.com/designsbysm/server-go/middleware"
	"github.com/gin-gonic/gin"
)

// TODO: remove, use token on front to test for login status

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/server")
	{
		group.GET("", middleware.AuthorizeJWT(), server)
	}
}
