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
