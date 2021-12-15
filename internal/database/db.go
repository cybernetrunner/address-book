package database

import (
	"fmt"
	"github.com/cyneruxyz/address-book/gen/proto"
	m "github.com/cyneruxyz/address-book/internal/database/model"
	"github.com/cyneruxyz/address-book/pkg/util"
	"gorm.io/gorm"
	"strings"
)

var (
	conf util.Config

	model = &m.Fields{}
	DSN   = fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBName,
		conf.DBPort,
		conf.DBSSLMode,
		conf.DBTimezone,
	)
)

type Database struct {
	ORM *gorm.DB
}

func (db *Database) CreateItem(field *proto.AddressField) error {
	return db.ORM.Create(model.Prepare(field)).Error
}

func (db *Database) ReadItem(param string) (fields []*proto.AddressField, err error) {
	var items []m.Fields

	param = convertWildcard(param)
	err = db.ORM.Model(model).
		Where("name LIKE ?", param).
		Or("address LIKE ?", param).
		Or("phone LIKE ?", param).
		Find(items).Error

	if err != nil {
		return nil, err
	}

	for _, item := range items {
		fields = append(fields, item.GetAddressField())
	}

	return fields, nil
}

func (db *Database) UpdateItem(phone *proto.Phone, replace *proto.AddressField) error {
	var item *m.Fields

	if err := db.ORM.Where("phone = ?", phone).First(item).Error; err != nil {
		return err
	}

	item.Name = replace.Name
	item.Address = replace.Address
	item.Phone = replace.Phone.Phone

	return db.ORM.Save(item).Error
}

func (db *Database) DeleteItem(phone *proto.Phone) {
	var item *m.Fields

	db.ORM.Where("phone = ?", phone).Delete(item)
}

func convertWildcard(s string) string {
	s = strings.ReplaceAll(s, "?", "_")
	return strings.ReplaceAll(s, "*", "%")
}
