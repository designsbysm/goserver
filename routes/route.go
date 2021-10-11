package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/spf13/viper"

	"github.com/designsbysm/ginmiddleware"
	"github.com/designsbysm/goserver/routes/api"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		ginmiddleware.Logger(),
		ginmiddleware.Error(),
	)

	group := r.Group("")
	{
		api.AddRoute(group)
		r.NoRoute(ServeClient)
	}
}

func ServeClient(c *gin.Context) {
	client := viper.GetString("client.server")
	u, err := url.Parse(client)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else if strings.HasPrefix(c.Request.RequestURI, "/api") {
		// do nothing
	} else {
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
