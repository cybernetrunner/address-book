package app

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/pkg/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/profclems/go-dotenv"
	"google.golang.org/grpc"

	"context"
	"net"
	"net/http"
)

const (
	errProtoHandlers  = "Fatal error generated proto handlers: %s "
	errListenAndServe = "Fatal error of http controller: %s "
)

func Run(conf *dotenv.DotEnv, db *database.Database) {
	serSvr := NewServer(db)
	grpcPort := conf.GetString("SERVER_GRPC_PORT")
	httpPort := conf.GetString("SERVER_HTTP_PORT")

	// Initialize context and defer canceling this context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Registering Service Handlers
	util.ErrorHandler(
		errProtoHandlers,
		proto.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts))

	util.ErrorHandler(
		errProtoHandlers,
		proto.RegisterAddressBookServiceHandlerServer(ctx, mux, serSvr))

	// Start gRPC server
	port, _ := net.Listen("tcp", grpcPort)
	srvRPC := grpc.NewServer()

	go func() {
		_ = srvRPC.Serve(port)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	util.ErrorHandler(
		errListenAndServe,
		http.ListenAndServe(httpPort, mux))
}
