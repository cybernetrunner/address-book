package main

import (
	// run "github.com/cyneruxyz/address-book/internal/address-book"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	"net"
)

const (
	StrStart = "Service started"
	StrStop  = "Service stopped"
	StrServe = "Serving gRPC on 0.0.0.0:8080"
)

func main() {
	log.Info(StrStart)
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	// helloworldpb.RegisterGreeterServer(s, &run.Server{})
	// Serve gRPC Server
	log.Println(StrServe)

	log.Fatalln(StrStop, s.Serve(lis))
}
