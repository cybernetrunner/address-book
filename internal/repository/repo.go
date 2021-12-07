package repository

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	"sync"
)

type phone = string

type AddressBook struct {
	sync.RWMutex
	Book map[phone]*proto.AddressField
}

func NewAddressBook() *AddressBook {
	return &AddressBook{}
}

func (ab *AddressBook) CreateAddressField(af *proto.AddressField) {
	ab.Lock()
	defer ab.Unlock()

	ab.Book[af.Phone] = af
}

func (ab *AddressBook) GetAddressFields(param string) (fields []*proto.AddressField) {
	ab.RLock()
	defer ab.RUnlock()

	if field, ok := ab.getAddressByPhone(param); ok {
		return []*proto.AddressField{field}
	}
	return ab.getAddressArray(param)
}

func (ab *AddressBook) UpdateAddressField(original, replace *proto.AddressField) (ok bool) {
	ab.Lock()
	defer ab.Unlock()

	if field, ok := ab.getAddressByPhone(original.Phone); ok {
		ab.Book[field.Phone] = replace
		return ok
	}
	return false
}

func (ab *AddressBook) DeleteAddressField(af *proto.AddressField) {
	ab.Lock()
	defer ab.Unlock()

	delete(ab.Book, af.Phone)
}

func (ab *AddressBook) getAddressByPhone(phone string) (field *proto.AddressField, ok bool) {
	if s, ok := ab.Book[phone]; ok {
		return s, ok
	}

	return nil, false
}

func (ab *AddressBook) getAddressArray(param string) (field []*proto.AddressField) {
	for k, v := range ab.Book {
		if param == v.Name || param == v.Address {
			field = append(field, ab.Book[k])
		}
	}
	return field
}
