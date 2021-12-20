package app

import (
	"context"
	"github.com/cyneruxyz/address-book/gen/proto"
)

const successMsg = "Procedure completed successfully"

type server struct {
	proto.AddressBookServiceServer
	repo Storage
}

func NewServer(repo Storage) server {
	return server{repo: repo}
}

func (s server) Echo(_ context.Context, request *proto.EchoRequest) (*proto.Response, error) {
	return &proto.Response{Message: request.Message}, nil
}

func (s server) Create(_ context.Context, request *proto.AddressFieldRequest) (*proto.Response, error) {
	if err := s.repo.CreateItem(request.Field); err != nil {
		return nil, err
	}

	return &proto.Response{Message: successMsg}, nil
}

func (s server) Read(_ context.Context, query *proto.AddressFieldQuery) (*proto.AddressFieldResponse, error) {
	return s.repo.ReadItem(query.Param)
}

func (s server) Update(_ context.Context, request *proto.AddressFieldUpdateRequest) (*proto.Response, error) {
	if err := s.repo.UpdateItem(request.Phone, request.ReplacementField); err != nil {
		return nil, err
	}

	return &proto.Response{Message: successMsg}, nil
}

func (s server) Delete(_ context.Context, phone *proto.Phone) (*proto.Response, error) {
	s.repo.DeleteItem(phone)

	return &proto.Response{Message: successMsg}, nil
}

func (s server) MustEmbedUnimplementedAddressBookServiceServer() {}
