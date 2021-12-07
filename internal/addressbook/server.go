package addressbook

import (
	"context"
	"errors"
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/repository"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	errStorageNotResponses = "Storage not response"
	grpcServerEndpoint     = "localhost:9090"
)

var repo = repository.NewAddressBook()

type server struct {
	proto.APIServiceServer
}

func (s *server) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.Response, error) {
	return &proto.Response{Message: req.Message}, nil
}

func (s *server) CreateAddressField(ctx context.Context, req *proto.AddressFieldRequest) (*proto.AddressField, error) {
	repo.CreateAddressField(req.Field)

	return req.Field, nil
}

func (s *server) ReadAddressField(ctx context.Context, query *proto.AddressFieldQuery) ([]*proto.AddressField, error) {
	return repo.GetAddressFields(query.Param), nil
}

func (s *server) UpdateAddressField(ctx context.Context, req *proto.AddressFieldUpdateRequest) (*proto.Response, error) {
	if ok := repo.UpdateAddressField(req.OriginalField, req.ReplacementField); !ok {
		//goland:noinspection GoErrorStringFormat
		return nil, errors.New(errStorageNotResponses)
	}

	return &proto.Response{Message: "Address field updated successfully"}, nil
}

func (s *server) DeleteAddressField(ctx context.Context, req *proto.AddressFieldRequest, opts ...grpc.CallOption) (*proto.Response, error) {
	repo.DeleteAddressField(req.Field)

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
