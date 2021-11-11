package api

import (
	"fmt"

	"github.com/designsbysm/server-go/client"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() error {
	router := gin.New()
	AddRoute(router)
	client.AddRoute(router)

	address := viper.GetString("api.address")
	tls := viper.GetBool("api.tls")

	if viper.GetBool("gin.release") {
		security := "HTTP"
		if tls {
			security = " (HTTPS)"
		}
		timber.Info(fmt.Sprintf("API: listening on %s%s", address, security))
	}

	if tls {
		if err := router.RunTLS(
			address,
			viper.GetString("ssl.cert"),
			viper.GetString("ssl.key"),
		); err != nil {
			return fmt.Errorf("API: %s", err.Error())
		}
	} else if err := router.Run(address); err != nil {
		return fmt.Errorf("API: %s", err.Error())
	}

	timber.Info("API: closing")

	return nil
}
