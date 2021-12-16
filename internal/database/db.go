package database

import (
	"fmt"
	"github.com/cyneruxyz/address-book/gen/proto"
	m "github.com/cyneruxyz/address-book/internal/database/model"
	"github.com/profclems/go-dotenv"
	"gorm.io/gorm"
	"strings"
)

var model = &m.Fields{}

func GetDSN(conf *dotenv.DotEnv) string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		conf.GetString("DB_HOST"),
		conf.GetString("DB_USER"),
		conf.GetString("DB_NAME"),
		conf.GetString("DB_PORT"),
		conf.GetString("DB_SSLMODE"),
		conf.GetString("DB_PASSWORD"),
	)
}

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
