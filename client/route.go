package client

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(r *gin.Engine) {
	client := viper.GetString("client.server")
	u, err := url.Parse(client)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)

	r.NoRoute(serve(proxy))
}

func serve(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
