package model

import (
	"github.com/cyneruxyz/address-book/gen/proto"
)

// Fields gorm.Model definition
type Fields struct {
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
	Phone   string `gorm:"unique, primaryKey"`
}

func (m *Fields) Prepare(af *proto.AddressField) *Fields {
	return &Fields{
		Name:    af.Name,
		Address: af.Address,
		Phone:   af.Phone.Phone,
	}
}

func (m *Fields) GetDTO() *proto.AddressField {
	return &proto.AddressField{
		Name:    m.Name,
		Address: m.Address,
		Phone: &proto.Phone{
			Phone: m.Phone,
		},
	}
}
