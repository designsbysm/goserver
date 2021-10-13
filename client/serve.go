package client

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func serve(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
