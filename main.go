package main

import (
	"fmt"
	"os"

	"github.com/designsbysm/goserver/database"
	"github.com/designsbysm/goserver/routes"

	"github.com/designsbysm/logger/v2"
	"github.com/designsbysm/loggerfile"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	// setup logger
	logger.New(
		os.Stdout,
		viper.GetInt("logger.cli.level"),
		viper.GetBool("logger.cli.colorful"),
		viper.GetBool("logger.cli.title"),
		viper.GetString("logger.cli.timestamp"),
	)

	f := loggerfile.New(
		viper.GetString("logger.file.path"),
	)
	logger.New(
		f,
		viper.GetInt("logger.file.level"),
		false,
		viper.GetBool("logger.file.title"),
		viper.GetString("logger.file.timestamp"),
	)

	// run the server
	router := gin.New()
	routes.AddRoute(router)
	port := viper.GetString("server.port")
	protocol := viper.GetString("server.protocol")

	if viper.GetBool("gin.release") {
		logger.Write(
			logger.LevelInfo,
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

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
