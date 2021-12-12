package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var (
	conf *viper.Viper
	db   *database.Database
	err  error

	errConfigFile = "Fatal error config file: %w \n"
	errDatabase   = "Fatal error database: %w \n"
)

func main() {
	conf = viper.New()
	conf.SetConfigFile(".env")

	if err = conf.ReadInConfig(); err != nil {
		glog.Fatalf(errConfigFile, err)
	}

	db, err = database.New(conf)
	if err != nil {
		glog.Fatalf(errDatabase, err)
	}

	if err = db.Migrate(); err != nil {
		glog.Fatalf(errDatabase, err)
	}

	if err = app.Run(conf, db); err != nil {
		glog.Fatal(err)
	}

	defer glog.Flush()
}
