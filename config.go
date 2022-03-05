package main

import (
	"errors"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func config() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./config.yaml not found")
		} else {
			return err
		}
	}

	// env variables
	viper.BindEnv("URL_FRONTEND")

	// cli flags
	address := viper.GetString("api.address")
	tls := viper.GetBool("api.tls")

	flag.StringVar(&address, "address", address, "api address")
	flag.BoolVar(&tls, "tls", tls, "use TLS")
	flag.Parse()

	viper.Set("api.address", address)
	viper.Set("api.tls", tls)

	// setup stuff
	if viper.GetBool("gin.release") {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}
