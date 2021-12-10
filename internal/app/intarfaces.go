package app

import (
	"github.com/cyneruxyz/address-book/gen/proto"
)

type Storage interface {
	CreateItem(field *proto.AddressField) error
	ReadItem(param string) ([]*proto.AddressField, error)
	UpdateItem(phone *proto.Phone, replace *proto.AddressField) error
	DeleteItem(phone *proto.Phone) error
}
