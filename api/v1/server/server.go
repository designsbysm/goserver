package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func server(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
}
