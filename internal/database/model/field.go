package model

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"gorm.io/gorm"
)

// AddressField gorm.Model definition
type AddressField struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
	Phone   string `gorm:"unique, primaryKey"`
}

func (model *AddressField) Prepare(af *proto.AddressField) *AddressField {
	return &AddressField{
		Name:    af.Name,
		Address: af.Address,
		Phone:   af.Phone.Phone,
	}
}

func (model *AddressField) GetAddressField() *proto.AddressField {
	return &proto.AddressField{
		Name:    model.Name,
		Address: model.Address,
		Phone: &proto.Phone{
			Phone: model.Phone,
		},
	}
}
