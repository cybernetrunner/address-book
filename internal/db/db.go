package db

import (
	"github.com/cyneruxyz/address-book/gen/proto"
	m "github.com/cyneruxyz/address-book/internal/db/model"
	"gorm.io/gorm"
	"strings"
)

var (
	model = &m.Fields{}
	DSN   = "host=0.0.0.0 user=gorm dbname=gorm port=5432 sslmode=disable password=gorm12345"
)

type Database struct {
	ORM *gorm.DB
}

func (db *Database) CreateItem(field *proto.AddressField) error {
	return db.ORM.Create(model.Prepare(field)).Error
}

func (db *Database) ReadItem(p string) (buf *proto.AddressFieldResponse, err error) {
	var items []m.Fields

	p = convertWildcard(p)
	err = db.ORM.Model(model).
		Where("name LIKE ?", p).
		Or("address LIKE ?", p).
		Or("phone LIKE ?", p).
		Find(items).Error

	if err != nil {
		return nil, err
	}

	for i, item := range items {
		buf.Fields[i] = item.GetDTO()
	}

	return buf, nil
}

func (db *Database) UpdateItem(p *proto.Phone, r *proto.AddressField) error {
	var item *m.Fields

	if err := db.ORM.Where("p = ?", p).First(item).Error; err != nil {
		return err
	}

	item.Name = r.Name
	item.Address = r.Address
	item.Phone = r.Phone.Phone

	return db.ORM.Save(item).Error
}

func (db *Database) DeleteItem(p *proto.Phone) {
	var item *m.Fields

	db.ORM.Where("phone = ?", p).Delete(item)
}

func convertWildcard(s string) string {
	s = strings.ReplaceAll(s, "?", "_")
	return strings.ReplaceAll(s, "*", "%")
}
