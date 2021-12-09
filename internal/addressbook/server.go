package addressbook

import (
	"context"
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/repository"
	"github.com/cyneruxyz/address-book/internal/repository/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	repo               repository.Repository = storage.NewAddressBook()
	grpcServerEndpoint                       = "localhost:9090"
)

type server struct {
	proto.APIServiceServer
}

func (s *server) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.Response, error) {
	return &proto.Response{Message: req.Message}, nil
}

func (s *server) CreateAddressField(ctx context.Context, req *proto.AddressFieldRequest) (*proto.AddressField, error) {
	if err := repo.CreateAddressField(req.Field); err != nil {
		return nil, err
	}

	return req.Field, nil
}

func (s *server) ReadAddressField(ctx context.Context, query *proto.AddressFieldQuery) ([]*proto.AddressField, error) {
	return repo.GetAddressFields(query.Param)
}

func (s *server) UpdateAddressField(ctx context.Context, req *proto.AddressFieldUpdateRequest) (*proto.Response, error) {
	if err := repo.UpdateAddressField(req.Phone, req.ReplacementField); err != nil {
		return nil, err
	}

	return &proto.Response{Message: "Address field updated successfully"}, nil
}

func (s *server) DeleteAddressField(ctx context.Context, req *proto.Phone, opts ...grpc.CallOption) (*proto.Response, error) {
	if err := repo.DeleteAddressField(req); err != nil {
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

	err := proto.RegisterAPIServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterAPIServiceHandlerServer(ctx, mux, server{}.APIServiceServer)
	if err != nil {
		return err
	}

	port, _ := net.Listen("tcp", grpcServerEndpoint)
	srv := grpc.NewServer()

	go func() {
		_ = srv.Serve(port)
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
