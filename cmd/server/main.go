package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/internal/database/model"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conf *viper.Viper
	db   *database.Database
	err  error

	errConfigFile = "Fatal error config file: %w \n"
	errDatabase   = "Fatal error database: %w \n"
)

func init() {
	// config initialization
	conf = viper.New()
	conf.SetConfigFile(".env")

	if err = conf.ReadInConfig(); err != nil {
		glog.Fatalf(errConfigFile, err)
	}

	// orm initialization
	orm, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  database.DSN,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)
	if err != nil {
		glog.Fatalf(errDatabase, err)
	}

	db = &database.Database{ORM: orm}

	//  migrate model to database
	if err = db.ORM.AutoMigrate(&model.Fields{}); err != nil {
		glog.Fatalf(errDatabase, err)
	}
}

func main() {
	if err = app.Run(conf, db); err != nil {
		glog.Fatal(err)
	}

	defer glog.Flush()
}
