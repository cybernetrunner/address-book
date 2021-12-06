package addressbook

import (
	"context"
	"errors"
	"flag"
	apipb "github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/repository"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	ErrDataNotFound        = "Data not found"
	ErrStorageNotResponses = "Storage not response"
)

var (
	// comand-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	repo               = repository.NewAddressBook()
)

type server struct {
	apipb.APIServiceServer
}

func (s *server) Echo(ctx context.Context, req *apipb.EchoRequest) (*apipb.Response, error) {
	return &apipb.Response{Message: req.Message}, nil
}

func (s *server) CreateAddressField(ctx context.Context, req *apipb.AddressFieldRequest) (*apipb.AddressField, error) {
	if ok := repo.CreateAddressField(req.Field); !ok {
		//goland:noinspection GoErrorStringFormat
		return nil, errors.New(ErrStorageNotResponses)
	}

	return req.Field, nil
}

func (s *server) ReadAddressField(ctx context.Context, query *apipb.AddressFieldQuery) ([]*apipb.AddressField, error) {
	res, ok := repo.GetAddressFields(query.Param)

	if !ok {
		//goland:noinspection GoErrorStringFormat
		return nil, errors.New(ErrDataNotFound)
	}

	return res, nil
}

func (s *server) UpdateAddressField(ctx context.Context, req *apipb.AddressFieldUpdateRequest) (*apipb.Response, error) {
	if ok := repo.UpdateAddressField(req.OriginalField, req.ReplacementField); !ok {
		//goland:noinspection GoErrorStringFormat
		return nil, errors.New(ErrStorageNotResponses)
	}

	return &apipb.Response{Message: "Address field updated successfully"}, nil
}

func (s *server) DeleteAddressField(ctx context.Context, req *apipb.AddressFieldRequest, opts ...grpc.CallOption) (*apipb.Response, error) {
	if ok := repo.DeleteAddressField(req.Field); !ok {
		//goland:noinspection GoErrorStringFormat
		return nil, errors.New(ErrStorageNotResponses)
	}

	return &apipb.Response{Message: "Address field deleted successfully"}, nil
}

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := apipb.RegisterAPIServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	err = apipb.RegisterAPIServiceHandlerServer(ctx, mux, server{}.APIServiceServer)
	if err != nil {
		return err
	}

	port, _ := net.Listen("tcp", *grpcServerEndpoint)
	srv := grpc.NewServer()

	go func() {
		err := srv.Serve(port)
		if err != nil {
			return
		}
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
