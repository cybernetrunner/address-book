package repository

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	_ "github.com/golang/mock/mockgen/model"
)

type Repository interface {
	CreateAddressField(field *proto.AddressField) error
	GetAddressFields(param string) ([]*proto.AddressField, error)
	UpdateAddressField(phone *proto.Phone, replace *proto.AddressField) error
	DeleteAddressField(phone *proto.Phone) error
}
