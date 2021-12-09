package storage

import (
	"github.com/cyneruxyz/address-book/gen/proto"

	"fmt"
	"regexp"
	"sync"
)

const elementByParameterNotFound = "element with parameter %s not found"

type AddressBook struct {
	sync.RWMutex
	Book map[*proto.Phone]*proto.AddressField
}

func New() *AddressBook {
	return &AddressBook{}
}

func (ab *AddressBook) CreateItem(af *proto.AddressField) error {
	ab.Lock()
	defer ab.Unlock()
	if _, ok := ab.Book[af.Phone]; ok {
		return fmt.Errorf("an item with the same phone number already exists")
	}

	ab.Book[af.Phone] = af
	return nil
}

func (ab *AddressBook) ReadItem(param string) (fields []*proto.AddressField, err error) {
	ab.RLock()
	defer ab.RUnlock()

	if field, ok := ab.getItemByPhone(&proto.Phone{Phone: param}); ok {
		return []*proto.AddressField{field}, nil
	}

	if field, ok := ab.getItemArray(param); ok {
		return field, nil
	}

	return nil, fmt.Errorf(elementByParameterNotFound, param)
}

func (ab *AddressBook) UpdateItem(phone *proto.Phone, replace *proto.AddressField) error {
	ab.Lock()
	defer ab.Unlock()
	if _, ok := ab.Book[phone]; ok {
		ab.Book[phone] = replace
		return nil
	}
	return fmt.Errorf(elementByParameterNotFound, phone)
}

func (ab *AddressBook) DeleteItem(p *proto.Phone) (err error) {
	ab.Lock()
	defer ab.Unlock()

	if _, ok := ab.Book[p]; ok {
		delete(ab.Book, p)
		return nil
	}
	return fmt.Errorf(elementByParameterNotFound, p)

}

func (ab *AddressBook) getItemByPhone(p *proto.Phone) (field *proto.AddressField, ok bool) {
	if s, ok := ab.Book[p]; ok {
		return s, ok
	}

	return nil, false
}

func (ab *AddressBook) getItemArray(param string) (field []*proto.AddressField, ok bool) {
	for k, v := range ab.Book {
		if wildcard(v.Name, param) || wildcard(v.Address, param) {
			field = append(field, ab.Book[k])
		}
	}

	if len(field) == 0 {
		return nil, false
	}

	return field, true
}

func wildcard(check, compare string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("%s.*", compare), check)

	return match
}
