package main

import (
	"bytes"
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/pkg/config"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var (
	conf *config.Config
	db   *database.Database
	err  error

	errConfigFile = "Fatal error config file: %w \n"
	errDatabase   = "Fatal error database: %w \n"
)

func main() {
	conf = viper.New()
	conf.SetConfigType("yaml")

	if err = conf.ReadConfig(bytes.NewBuffer(config.Yaml)); err != nil {
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
