package v1

import (
	"github.com/designsbysm/goserver/routes/api/v1/session"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		session.AddRoute(group)
	}
}
