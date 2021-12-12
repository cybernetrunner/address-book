package database

import (
	"fmt"
	"github.com/cyneruxyz/address-book/gen/proto"
	"github.com/cyneruxyz/address-book/internal/database/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var fieldModel = &model.AddressField{}

type Database struct {
	orm *gorm.DB
}

func New(conf *viper.Viper) (*Database, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.GetString("DB_USER"),
		conf.GetString("DB_PASSWORD"),
		conf.GetString("DB_NAME"),
		conf.GetString("DB_PORT"),
		conf.GetString("DB_SSLMODE"),
		conf.GetString("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{},
	)

	return &Database{db}, err
}

func (db *Database) Migrate() error {
	return db.orm.AutoMigrate(fieldModel)
}

func (db *Database) CreateItem(field *proto.AddressField) error {
	return db.orm.Create(fieldModel.Prepare(field)).Error
}

func (db *Database) ReadItem(param string) (fields []*proto.AddressField, err error) {
	var items []model.AddressField

	err = db.orm.Model(fieldModel).
		Where("name = ?", param).
		Or("address = ?", param).
		Or("phone = ?", param).
		Limit(100).Find(items).Error

	if err != nil {
		return nil, err
	}

	for _, item := range items {
		fields = append(fields, item.GetAddressField())
	}

	return fields, nil
}

func (db *Database) UpdateItem(phone *proto.Phone, replace *proto.AddressField) error {
	var item *model.AddressField

	if err := db.orm.Where("phone = ?", phone).First(item).Error; err != nil {
		return err
	}

	item.Name = replace.Name
	item.Address = replace.Address
	item.Phone = replace.Phone.Phone

	return db.orm.Save(item).Error
}

func (db *Database) DeleteItem(phone *proto.Phone) error {
	var item *model.AddressField

	if err := db.orm.Where("phone = ?", phone).First(item).Error; err != nil {
		return err
	}

	if err := db.orm.Where("phone = ?", phone).Delete(item).Error; err != nil {
		return err
	}

	return nil
}
