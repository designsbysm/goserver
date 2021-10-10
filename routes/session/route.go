package session

import (
	"github.com/designsbysm/goserver/middleware"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/session")
	{
		group.POST("/login", login)
		group.POST("/logout", middleware.AuthorizeJWT(), logout)
	}
}
