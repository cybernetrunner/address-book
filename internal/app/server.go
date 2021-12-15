package app

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/pkg/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"context"
	"net"
	"net/http"
)

var ()

func Run(conf util.Config, db *database.Database) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, conf.ServerGRPCPort, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterAddressBookServiceHandlerServer(ctx, mux, server{repo: db}.AddressBookServiceServer)
	if err != nil {
		return err
	}

	port, _ := net.Listen("tcp", conf.ServerGRPCPort)
	srv := grpc.NewServer()

	go func() {
		_ = srv.Serve(port)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(conf.ServerHTTPPort, mux)
}
