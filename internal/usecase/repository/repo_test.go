package repository_test

import (
	"address-book/internal/entity"
	"address-book/internal/usecase/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	name    = "John Jr. Smith"
	address = "Baker str 221"
	phone   = "88005553535"

	stubRepo  = repository.NewRepository()
	stubModel = &entity.AddressField{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
)

func TestRepository_GetItem(t *testing.T) {
	stubRepo := stubRepo.AddField(name, address, phone)
	item := stubRepo.GetItem(0)

	assert.NotNil(t, *item)
	assert.Equal(t, stubModel, item)
}

func TestRepository_AddField(t *testing.T) {
	repo := repository.NewRepository()
	exampleRepo := append(*repo, stubModel)
	stubRepo := stubRepo.AddField(name, address, phone)

	assert.NotNil(t, stubRepo)
	assert.Equal(t, exampleRepo, stubRepo)
}

func TestRepository_DeleteField(t *testing.T) {
	stubRepo := stubRepo.AddField(name, address, phone)

	if index, ok := stubRepo.FindField(phone); ok {
		stubRepo = stubRepo.DeleteField(index)
		_, ok = stubRepo.FindField(phone)

		assert.Equal(t, false, ok)
	} else {
		t.Fatal("FindField not worked")
	}

}
