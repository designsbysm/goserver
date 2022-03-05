package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	router := gin.New()
	AddRoute(router)
	// client.AddRoute(router)

	address := viper.GetString("api.address")
	tls := viper.GetBool("api.tls")

	// run
	go func() {
		if tls {
			if err := router.RunTLS(
				address,
				viper.GetString("ssl.cert"),
				viper.GetString("ssl.key"),
			); err != nil {
				timber.Error("API:", err)
			}
		} else if err := router.Run(address); err != nil {
			timber.Error("API:", err)
		}
	}()

	// notify
	if viper.GetBool("gin.release") {
		security := "HTTP"
		if tls {
			security = "HTTPS"
		}

		timber.Info(fmt.Sprintf("API: listening on %s (%s)", address, security))
	}

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// close
	fmt.Println("")
	timber.Info("API: closed")
}
