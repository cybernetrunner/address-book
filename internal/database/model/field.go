package model

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"gorm.io/gorm"
)

// Fields gorm.Model definition
type Fields struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
	Phone   string `gorm:"unique, primaryKey"`
}

func (model *Fields) Prepare(af *proto.AddressField) *Fields {
	return &Fields{
		Name:    af.Name,
		Address: af.Address,
		Phone:   af.Phone.Phone,
	}
}

func (model *Fields) GetAddressField() *proto.AddressField {
	return &proto.AddressField{
		Name:    model.Name,
		Address: model.Address,
		Phone: &proto.Phone{
			Phone: model.Phone,
		},
	}
}
