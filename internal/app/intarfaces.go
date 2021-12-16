package app

import "github.com/cyneruxyz/address-book/gen/proto"

type Storage interface {
	CreateItem(*proto.AddressField) error
	ReadItem(param string) (*proto.AddressFieldResponse, error)
	UpdateItem(*proto.Phone, *proto.AddressField) error
	DeleteItem(*proto.Phone)
}
