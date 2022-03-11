package config

import (
	"errors"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Environment() error {
	// load config.yaml
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./config.yaml not found")
		} else {
			return err
		}
	}

	// load .env
	file, err := os.Open("../.env")
	if err == nil {
		defer file.Close()

		viper.SetConfigType("env")
		viper.MergeConfig(file)
	}

	// env variables
	viper.BindEnv("URL_FRONTEND")
	viper.BindEnv("URL_DATABASE")

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
