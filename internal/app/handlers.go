package app

import (
	"context"
	"github.com/cyneruxyz/address-book/gen/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.AddressBookServiceServer
	repo Storage
}

func (s *server) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.Response, error) {
	return &proto.Response{Message: req.Message}, nil
}

func (s *server) Create(ctx context.Context, req *proto.AddressFieldRequest) (*proto.AddressField, error) {
	if err := s.repo.CreateItem(req.Field); err != nil {
		return nil, err
	}

	return req.Field, nil
}

func (s *server) Read(ctx context.Context, query *proto.AddressFieldQuery) ([]*proto.AddressField, error) {
	return s.repo.ReadItem(query.Param)
}

func (s *server) Update(ctx context.Context, req *proto.AddressFieldUpdateRequest) (*proto.Response, error) {
	if err := s.repo.UpdateItem(req.Phone, req.ReplacementField); err != nil {
		return nil, err
	}

	return &proto.Response{Message: "Address field updated successfully"}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.Phone, opts ...grpc.CallOption) (*proto.Response, error) {
	s.repo.DeleteItem(req)

	return &proto.Response{Message: "Address field deleted successfully"}, nil
}
