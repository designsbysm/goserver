package main

import (
	"fmt"

	"github.com/designsbysm/server-go/api"
	"github.com/designsbysm/server-go/client"

	"github.com/designsbysm/logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func server() error {
	router := gin.New()
	api.AddRoute(router)
	client.AddRoute(router)

	port := viper.GetString("server.port")
	protocol := viper.GetString("server.protocol")

	if viper.GetBool("gin.release") {
		logger.Info(
			fmt.Sprintf("serving %s over %s", protocol, port),
		)
	}

	if protocol == "HTTPS" {
		if err := router.RunTLS(
			port,
			viper.GetString("server.https.cert"),
			viper.GetString("server.https.key"),
		); err != nil {
			panic(err)
		}
	}

	return router.Run(port)
}
