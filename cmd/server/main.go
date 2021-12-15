package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/pkg/util"
	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conf util.Config
	db   *database.Database
	err  error

	errConfigFile = "Fatal error config file: %w \n"
	errDatabase   = "Fatal error database: %w \n"
)

func init() {
	// util initialization
	conf, err = util.LoadConfig(".")
	if err != nil {
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
}

func main() {
	if err = app.Run(conf, db); err != nil {
		glog.Fatal(err)
	}

	defer glog.Flush()
}
