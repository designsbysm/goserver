package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func config() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./default.yaml not found")
		} else {
			return err
		}
	}

	viper.SetConfigName("override")
	viper.AddConfigPath(".")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./server.yaml not found")
		} else {
			return err
		}
	}

	// cli flags
	var port int
	var https bool

	flag.IntVar(&port, "port", 0, "server port")
	flag.BoolVar(&https, "https", false, "user HTTPS")
	flag.Parse()

	if port > 0 {
		viper.Set("server.port", fmt.Sprintf(":%d", port))
	}

	if https {
		viper.Set("server.protocol", "HTTPS")
	}

	// setup stuff
	if viper.GetBool("gin.release") {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}
