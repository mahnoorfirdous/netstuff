package main

import (
	"context"
	"flag"
	"grsidecar/pbgen"
	"grsidecar/sample"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dialaddress := flag.String("address", "", "we are dialing this address")
	flag.Parse()
	log.Infof("Dialing server %s", *dialaddress)

	connection, err := grpc.Dial(*dialaddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Cannot dial server! Error: ", err)
	}

	AlertClient := pbgen.NewCaterAlertRequestClient(connection)
	alert := sample.NewAlert()

	AlertClient.CaterAlert(context.Background(), alert)

}