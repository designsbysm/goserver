package client

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

func serve(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api/") {
			c.Status(http.StatusNotFound)
			return
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
