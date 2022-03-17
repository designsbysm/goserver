package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func addCORS(r *gin.Engine) {
	client := viper.GetString("URL_FRONTEND")

	if client == "" {
		return
	}

	r.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins:     []string{client},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Length", "Content-Type", "Host", "Referrer", "Origin", "User-Agent"},
		AllowMethods:     []string{"DELETE", "GET", "POST", "OPTIONS", "PUT"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))
}
