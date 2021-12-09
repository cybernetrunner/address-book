package app

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"context"
	"net"
	"net/http"
)

var (
	repo     Storage = storage.New()
	httpPort         = "8081"
	grpcPort         = "9090"
)

type server struct {
	proto.AddressBookServiceServer
}

func (s *server) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.Response, error) {
	return &proto.Response{Message: req.Message}, nil
}

func (s *server) Create(ctx context.Context, req *proto.AddressFieldRequest) (*proto.AddressField, error) {
	if err := repo.CreateItem(req.Field); err != nil {
		return nil, err
	}

	return req.Field, nil
}

func (s *server) Read(ctx context.Context, query *proto.AddressFieldQuery) ([]*proto.AddressField, error) {
	return repo.ReadItem(query.Param)
}

func (s *server) Update(ctx context.Context, req *proto.AddressFieldUpdateRequest) (*proto.Response, error) {
	if err := repo.UpdateItem(req.Phone, req.ReplacementField); err != nil {
		return nil, err
	}

	return &proto.Response{Message: "Address field updated successfully"}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.Phone, opts ...grpc.CallOption) (*proto.Response, error) {
	if err := repo.DeleteItem(req); err != nil {
		return nil, err
	}

	return &proto.Response{Message: "Address field deleted successfully"}, nil
}

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterAddressBookServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterAddressBookServiceHandlerServer(ctx, mux, server{}.AddressBookServiceServer)
	if err != nil {
		return err
	}

	port, _ := net.Listen("tcp", grpcPort)
	srv := grpc.NewServer()

	go func() {
		_ = srv.Serve(port)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(httpPort, mux)
}
