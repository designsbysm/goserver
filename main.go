// TODO: add documentation
package main

import (
	"github.com/designsbysm/server-go/database"
)

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	if err := loggers(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := server(); err != nil {
		panic(err)
	}
}
