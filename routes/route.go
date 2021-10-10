package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/designsbysm/goserver/routes/session"

	"github.com/designsbysm/ginmiddleware"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		ginmiddleware.Logger(),
		ginmiddleware.Error(),
	)

	group := r.Group("/v1")
	{
		session.AddRoute(group)
	}

	r.NoRoute(ServeClient)
}

func ServeClient(c *gin.Context) {
	client := os.Getenv("CLIENT_SERVER")
	u, err := url.Parse(client)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else if strings.HasPrefix(c.Request.RequestURI, "/v1") {
		// do nothing
	} else if c.Request.URL.String() == "/favicon.ico" {
		c.Status(http.StatusNoContent)
	} else {
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
