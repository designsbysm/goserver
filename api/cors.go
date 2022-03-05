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
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))
}
