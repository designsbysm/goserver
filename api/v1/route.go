package v1

import (
	"github.com/designsbysm/server-go/api/v1/server"
	"github.com/designsbysm/server-go/api/v1/session"
	"github.com/designsbysm/server-go/api/v1/user"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		session.AddRoute(group)
		server.AddRoute(group)
		user.AddRoute(group)
	}
}
