package storage_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/app"
	mockapp "github.com/cyneruxyz/address-book/internal/mock"
	"github.com/cyneruxyz/address-book/internal/storage"
	"github.com/golang/mock/gomock"

	"testing"
)

var r app.Storage = storage.New()

func randomField() proto.AddressField {
	return proto.AddressField{
		Name:    gofakeit.Name(),
		Address: gofakeit.Street(),
		Phone: &proto.Phone{
			Phone: gofakeit.Phone(),
		},
	}
}

func TestCreteItem(t *testing.T) {
	address := randomField()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockapp.NewMockStorage(ctrl)

	repo.EXPECT().
		CreateItem(gomock.Eq(&address)).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateItem(&address); err != nil {
		t.Fail()
	}
}

func TestReadItem(t *testing.T) {
	address := randomField()
	param := address.Name

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockapp.NewMockStorage(ctrl)

	repo.EXPECT().
		ReadItem(gomock.Eq(param)).Times(1).
		Return(
			gomock.Eq(append([]*proto.AddressField{}, &address)),
			gomock.Nil(),
		)

	if err := r.CreateItem(&address); err != nil {
		t.Fail()
	}

	if arr, err := r.ReadItem(param); err != nil {
		t.Fail()
	} else if arr[0] != &address {
		t.Fail()
	}
}

func TestUpdateItem(t *testing.T) {
	addressOriginal := randomField()
	addressReplacer := randomField()
	phone := addressOriginal.Phone

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockapp.NewMockStorage(ctrl)

	repo.EXPECT().
		CreateItem(gomock.Eq(&addressOriginal)).
		Return(gomock.Nil()).Times(1)

	repo.EXPECT().
		UpdateItem(
			gomock.Eq(phone),
			gomock.Eq(&addressReplacer),
		).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateItem(&addressOriginal); err != nil {
		t.Fail()
	}

	if err := r.UpdateItem(phone, &addressReplacer); err != nil {
		t.Fail()
	}
}

func TestDeleteItem(t *testing.T) {
	address := randomField()
	phone := address.Phone

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockapp.NewMockStorage(ctrl)

	repo.EXPECT().
		DeleteItem(gomock.Eq(phone)).
		Return(gomock.Nil()).Times(1)

	if err := r.CreateItem(&address); err != nil {
		t.Fail()
	}

	if err := r.DeleteItem(phone); err != nil {
		t.Fail()
	}
}
