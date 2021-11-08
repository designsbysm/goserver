package rpc

import (
	"fmt"
	"net"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func Serve() error {
	port := viper.GetString("rpc.port")
	protocol := viper.GetString("rpc.protocol")

	_, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("RPC: %s", err.Error())
	}

	timber.Info(fmt.Sprintf("RPC: serving %s on %s", protocol, port))

	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	timber.Info("RPC: closing")

	return nil
}
