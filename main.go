// TODO: add documentation
package main

import (
	"github.com/designsbysm/server-go/api"
	"github.com/designsbysm/server-go/database"
)

func main() {
	if err := config(); err != nil {
		panic(err)
	}

	if err := loggers(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	// run each server
	api.Serve()
}
