package repository

import (
	"address-book/internal/entity"
)

type Repository []*entity.AddressField

func NewRepository() *Repository {
	return new(Repository)
}

func (r Repository) FindField(param string) (index int, ok bool) {
	for i, field := range r {
		if field.Name == param {
			return i, true
		}
	}
	for i, field := range r {
		if field.Phone == param {
			return i, true
		}
	}
	for i, field := range r {
		if field.Phone == param {
			return i, true
		}
	}

	return 0, false
}

func (r Repository) AddField(name, address, phone string) Repository {
	return append(r, entity.NewAddressField(name, address, phone))
}

func (r Repository) GetItem(index int) *entity.AddressField {
	return r[index]
}

func (r Repository) DeleteField(index int) Repository {
	return append(r[:index], r[index+1:]...)
}
