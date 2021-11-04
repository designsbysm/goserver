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

	port := viper.GetString("api.port")
	protocol := viper.GetString("api.protocol")

	if viper.GetBool("gin.release") {
		timber.Info(fmt.Sprintf("serving API on %s", port))
	}

	if protocol == "HTTPS" {
		if err := router.RunTLS(
			port,
			viper.GetString("ssl.cert"),
			viper.GetString("ssl.key"),
		); err != nil {
			return fmt.Errorf("API Server %s", err.Error())
		}
	}

	if err := router.Run(port); err != nil {
		return fmt.Errorf("API Server %s", err.Error())
	}

	timber.Info("closing API")

	return nil
}
