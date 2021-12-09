package repository_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/repository"
	"github.com/cyneruxyz/address-book/internal/repository/storage"
	"github.com/cyneruxyz/address-book/test/mock_repository"
	"github.com/golang/mock/gomock"
	"testing"
)

var r repository.Repository = storage.NewAddressBook()

func randomField() proto.AddressField {
	return proto.AddressField{
		Name:    gofakeit.Name(),
		Address: gofakeit.Street(),
		Phone: &proto.Phone{
			Phone: gofakeit.Phone(),
		},
	}
}

func TestCreteAddressField(t *testing.T) {
	address := randomField()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockRepository(ctrl)
	repo.EXPECT().
		CreateAddressField(gomock.Eq(&address)).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateAddressField(&address); err != nil {
		t.Fail()
	}
}

func TestGetAddressField(t *testing.T) {
	address := randomField()
	param := address.Name

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockRepository(ctrl)
	repo.EXPECT().
		GetAddressFields(gomock.Eq(param)).Times(1).
		Return(
			gomock.Eq(append([]*proto.AddressField{}, &address)),
			gomock.Nil(),
		)

	if err := r.CreateAddressField(&address); err != nil {
		t.Fail()
	}

	if arr, err := r.GetAddressFields(param); err != nil {
		t.Fail()
	} else if arr[0] != &address {
		t.Fail()
	}
}

func TestUpdateAddressField(t *testing.T) {
	addressOriginal := randomField()
	addressReplacer := randomField()
	phone := addressOriginal.Phone

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockRepository(ctrl)

	repo.EXPECT().
		CreateAddressField(gomock.Eq(&addressOriginal)).
		Return(gomock.Nil()).Times(1)

	repo.EXPECT().
		UpdateAddressField(
			gomock.Eq(phone),
			gomock.Eq(&addressReplacer),
		).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateAddressField(&addressOriginal); err != nil {
		t.Fail()
	}

	if err := r.UpdateAddressField(phone, &addressReplacer); err != nil {
		t.Fail()
	}
}

func TestDeleteAddressField(t *testing.T) {
	address := randomField()
	phone := address.Phone

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockRepository(ctrl)
	repo.EXPECT().
		DeleteAddressField(gomock.Eq(phone)).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateAddressField(&address); err != nil {
		t.Fail()
	}

	if err := r.DeleteAddressField(phone); err != nil {
		t.Fail()
	}
}
