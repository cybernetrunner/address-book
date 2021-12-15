package app

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/profclems/go-dotenv"
	"google.golang.org/grpc"

	"context"
	"net"
	"net/http"
)

var ()

func Run(conf *dotenv.DotEnv, db *database.Database) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, conf.GetString("SERVER_GRPC_PORT"), opts)
	if err != nil {
		return err
	}

	err = proto.RegisterAddressBookServiceHandlerServer(ctx, mux, server{repo: db}.AddressBookServiceServer)
	if err != nil {
		return err
	}

	port, _ := net.Listen("tcp", conf.GetString("SERVER_GRPC_PORT"))
	srv := grpc.NewServer()

	go func() {
		_ = srv.Serve(port)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(conf.GetString("SERVER_HTTP_PORT"), mux)
}
