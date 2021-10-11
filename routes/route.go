package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/spf13/viper"

	"github.com/designsbysm/ginmiddleware"
	"github.com/designsbysm/server-go/routes/api"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		ginmiddleware.Logger(),
		ginmiddleware.Error(),
	)

	// setup frontend client
	client := viper.GetString("client.server")
	u, err := url.Parse(client)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)

	group := r.Group("")
	{
		api.AddRoute(group)
		r.NoRoute(serveClient(proxy))
	}
}
