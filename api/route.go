package api

import (
	"github.com/designsbysm/ginmiddleware"
	v1 "github.com/designsbysm/server-go/api/v1"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		ginmiddleware.Logger(),
		ginmiddleware.Error(),
	)

	group := r.Group("api")
	{
		v1.AddRoute(group)
	}
}
