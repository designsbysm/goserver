package routes

import (
	"fmt"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func serveClient(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("proxy", c.Param("path"))
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
