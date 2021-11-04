package rpc

import (
	"fmt"
	"net"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func Serve() error {
	port := viper.GetString("rpc.port")

	_, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("RPC Server %s", err.Error())
	}

	timber.Info(fmt.Sprintf("serving RPC on %s", port))

	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	timber.Info("closing RPC")

	return nil
}
