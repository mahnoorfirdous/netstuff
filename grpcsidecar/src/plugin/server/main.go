package main

import (
	"flag"
	"grsidecar/pbgen"
	"grsidecar/service"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartAlertServer(port string) error {
	alertServer := service.NewAlertCaterServer(&service.AlertStore{})

	grpcServer := grpc.NewServer()
	pbgen.RegisterCaterAlertRequestServer(grpcServer, alertServer)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Info("Starting GRPC Listener...")
	log.Fatal(grpcServer.Serve(listener))
	return nil
}

func main() {

	serveon := flag.String("port", "0", "the port on which the server will listen")
	StartAlertServer(*serveon)
}
