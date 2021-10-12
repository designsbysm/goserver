package server

import (
	"github.com/designsbysm/server-go/middleware"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/server")
	{
		group.GET("/", middleware.AuthorizeJWT(), server)
	}
}
