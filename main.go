// TODO: add documentation
package main

import (
	"github.com/designsbysm/server-go/api"
	"github.com/designsbysm/server-go/database"
	"github.com/designsbysm/timber/v2"
	"golang.org/x/sync/errgroup"
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

	// run each server type
	eg := new(errgroup.Group)

	eg.Go(func() error { return api.Serve() })
	// eg.Go(func() error { return rpc.Serve() })

	err := eg.Wait()
	if err != nil {
		timber.Error(err)
	}
}
