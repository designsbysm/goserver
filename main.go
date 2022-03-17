// TODO: remove gin? https://echo.labstack.com
// TODO: add documentation
package main

import (
	"github.com/designsbysm/server-go/api"
	"github.com/designsbysm/server-go/config"
	"github.com/designsbysm/server-go/database"
)

func main() {
	if err := config.Environment(); err != nil {
		panic(err)
	}

	if err := config.Loggers(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	api.Serve()
}
