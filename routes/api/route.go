package api

import (
	v1 "github.com/designsbysm/goserver/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/api")
	{
		v1.AddRoute(group)
	}
}
